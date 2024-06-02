package orchestrator

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Task struct {
	ID            string  `json:"id"`
	Arg1          float64 `json:"arg1"`
	Arg2          float64 `json:"arg2"`
	Operation     string  `json:"operation"`
	OperationTime int     `json:"operation_time"`
}

var (
	tasks     = make(map[string]Task)
	taskQueue = make(chan Task, 100)
	taskMu    sync.Mutex
)

func getTask(w http.ResponseWriter, r *http.Request) {
	select {
	case task := <-taskQueue:
		json.NewEncoder(w).Encode(map[string]Task{"task": task})
	default:
		http.Error(w, "no task available", http.StatusNotFound)
	}
}

func submitTask(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ID     string  `json:"id"`
		Result float64 `json:"result"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	taskMu.Lock()
	taskResults[req.ID] = req.Result
	taskMu.Unlock()

	expressionTasksMu.Lock()
	defer expressionTasksMu.Unlock()
	for expressionID, tasks := range expressionTasks {
		for _, taskID := range tasks {
			if taskID == req.ID {
				// Проверяем, все ли задачи для данного выражения уже выполнены
				allTasksCompleted := true
				for _, tID := range tasks {
					if _, ok := taskResults[tID]; !ok {
						allTasksCompleted = false
						break
					}
				}

				if allTasksCompleted {
					// Обновляем результаты выражения и статус
					expression := expressions[expressionID]
					expression.Result = req.Result
					expression.Status = "completed"
					expressions[expressionID] = expression
				}

				return
			}
		}
	}
	w.WriteHeader(http.StatusOK)
}

func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
