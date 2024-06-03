package main

import (
	"calculator/internal/agent"
	"log"
	"os"
	"strconv"
)

func main() {

	agentCount, err := strconv.Atoi(os.Getenv("COMPUTING_POWER"))
	if err != nil {
		log.Fatalf("Ошибка загрузки переменной среды COMPUTING_POWER: %v", err)
	}

	for i := 0; i < agentCount; i++ {
		go agent.Run()
	}
	select {}
}
