package orchestrator

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// parseExpression разбирает арифметическое выражение на отдельные задачи
func parseExpression(expression string) []Task {
	var tasks []Task

	// Разделение выражения на отдельные операции
	operations := strings.Split(expression, " ")
	log.Println(operations)
	arg1, _ := strconv.ParseFloat(operations[0], 64)
	arg2, _ := strconv.ParseFloat(operations[2], 64)
	operation := operations[1]
	tasks = append(tasks, Task{
		ID:            generateID(),
		Arg1:          arg1,
		Arg2:          arg2,
		Operation:     operation,
		OperationTime: getOperationTime(operation), // Здесь нужно определить время выполнения операции
	})

	return tasks
}

func getOperationTime(operation string) int {
	var envVar string
	switch operation {
	case "+":
		envVar = "TIME_ADDITION_MS"
	case "-":
		envVar = "TIME_SUBTRACTION_MS"
	case "*":
		envVar = "TIME_MULTIPLICATIONS_MS"
	case "/":
		envVar = "TIME_DIVISIONS_MS"
	default:
		log.Println("Неизвестная операция:", operation)
		return 0
	}

	timeStr := os.Getenv(envVar)
	if timeStr == "" {
		log.Printf("Переменная среды %s не установлена\n", envVar)
		return 0
	}

	timeMS, err := strconv.Atoi(timeStr)
	if err != nil {
		log.Printf("Ошибка преобразования переменной среды %s: %v\n", envVar, err)
		return 0
	}

	return timeMS
}
