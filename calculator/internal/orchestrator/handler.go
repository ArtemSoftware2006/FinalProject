package orchestrator

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type Expression struct {
	ID     string  `json:"id"`
	Status string  `json:"status"`
	Result float64 `json:"result"`
}

var (
	expressions       = make(map[string]Expression)
	taskResults       = make(map[string]float64)
	expressionTasks   = make(map[string][]string) // Сопоставление идентификатора выражения с идентификаторами задач
	expressionTasksMu sync.Mutex
)

func RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/api/v1/calculate", addExpression).Methods("POST")
	r.HandleFunc("/api/v1/expressions", listExpressions).Methods("GET")
	r.HandleFunc("/api/v1/expressions/{id}", getExpression).Methods("GET")
	r.HandleFunc("/internal/task", getTask).Methods("GET")
	r.HandleFunc("/internal/task", submitTask).Methods("POST")
}

func addExpression(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Expression string `json:"expression"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tasks := parseExpression(req.Expression)
	expressionID := generateID()

	// Сохраняем идентификаторы задач для данного выражения
	expressionTasksMu.Lock()
	expressionTasks[expressionID] = make([]string, len(tasks))
	for i, task := range tasks {
		taskQueue <- task
		expressionTasks[expressionID][i] = task.ID
	}
	expressions[expressionID] = Expression{ID: expressionID, Status: "pending"}
	expressionTasksMu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": expressionID})
}

func listExpressions(w http.ResponseWriter, r *http.Request) {
	expressionTasksMu.Lock()
	defer expressionTasksMu.Unlock()
	var list []Expression
	for _, exp := range expressions {
		list = append(list, exp)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"expressions": list})
}

func getExpression(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	expressionTasksMu.Lock()
	exp, ok := expressions[id]
	expressionTasksMu.Unlock()
	if !ok {
		http.Error(w, "expression not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]Expression{"expression": exp})
}
