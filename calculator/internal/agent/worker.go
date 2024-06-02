package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Task struct {
	ID            string  `json:"id"`
	Arg1          float64 `json:"arg1"`
	Arg2          float64 `json:"arg2"`
	Operation     string  `json:"operation"`
	OperationTime int     `json:"operation_time"`
}

func Run() {
	for {
		task, err := getTask()
		if err != nil {
			log.Println("No tasks available")
			time.Sleep(2 * time.Second)
			continue
		}
		log.Println(task)
		result := performTask(task)
		log.Println("Result: ", result)
		submitResult(task.ID, result)
	}
}

func getTask() (Task, error) {
	resp, err := http.Get("http://localhost:8080/internal/task")
	if err != nil {
		return Task{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return Task{}, fmt.Errorf("no tasks available")
	}
	var respBody struct {
		Task Task `json:"task"`
	}
	json.NewDecoder(resp.Body).Decode(&respBody)
	return respBody.Task, nil
}

func performTask(task Task) float64 {
	time.Sleep(time.Duration(task.OperationTime) * time.Millisecond)
	switch task.Operation {
	case "+":
		return task.Arg1 + task.Arg2
	case "-":
		return task.Arg1 - task.Arg2
	case "*":
		return task.Arg1 * task.Arg2
	case "/":
		if task.Arg2 == 0 {
			return 0
		}
		return task.Arg1 / task.Arg2
	default:
		return 0
	}
}

func submitResult(id string, result float64) {
	data, _ := json.Marshal(map[string]interface{}{
		"id":     id,
		"result": result,
	})
	http.Post("http://localhost:8080/internal/task", "application/json", bytes.NewBuffer(data))
}
