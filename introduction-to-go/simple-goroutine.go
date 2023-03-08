package main

import (
    "fmt"
    "time"
)

func printHelloWorld() {
    fmt.Println("Hello world from goroutine!")
}

func main() {
    go printHelloWorld()
    fmt.Println("Hello from main!")
    time.Sleep(time.Second)
}
