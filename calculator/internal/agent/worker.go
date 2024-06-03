package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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
	timeAdditionMS        int
	timeSubtractionMS     int
	timeMultiplicationsMS int
	timeDivisionsMS       int
)

func Run() {
	loadEnvValues()
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
	var delay int
	switch task.Operation {
	case "+":
		delay = timeAdditionMS
	case "-":
		delay = timeSubtractionMS
	case "*":
		delay = timeMultiplicationsMS
	case "/":
		delay = timeDivisionsMS
	default:
		delay = 0
	}

	time.Sleep(time.Duration(delay) * time.Millisecond)
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
func loadEnvValues() {
	var err error
	timeAdditionMS, err = strconv.Atoi(os.Getenv("TIME_ADDITION_MS"))
	if err != nil {
		log.Fatalf("Ошибка загрузки переменной среды TIME_ADDITION_MS: %v", err)
	}

	timeSubtractionMS, err = strconv.Atoi(os.Getenv("TIME_SUBTRACTION_MS"))
	if err != nil {
		log.Fatalf("Ошибка загрузки переменной среды TIME_SUBTRACTION_MS: %v", err)
	}

	timeMultiplicationsMS, err = strconv.Atoi(os.Getenv("TIME_MULTIPLICATIONS_MS"))
	if err != nil {
		log.Fatalf("Ошибка загрузки переменной среды TIME_MULTIPLICATIONS_MS: %v", err)
	}

	timeDivisionsMS, err = strconv.Atoi(os.Getenv("TIME_DIVISIONS_MS"))
	if err != nil {
		log.Fatalf("Ошибка загрузки переменной среды TIME_DIVISIONS_MS: %v", err)
	}
}
