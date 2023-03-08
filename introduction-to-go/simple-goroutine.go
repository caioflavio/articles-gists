package main

import (
	"fmt"
	"time"
)

func count(id string) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Goroutine", id, "Tarefa:", i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go count("Worker 1")
	go count("Worker 2")
	go count("Worker 2")
	time.Sleep(time.Second * 3)
	fmt.Println("Fim da função main.")
}
