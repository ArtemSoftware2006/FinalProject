package orchestrator

import (
	"log"
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
		OperationTime: 200, // getOperationTime(operation), // Здесь нужно определить время выполнения операции
	})

	return tasks
}
