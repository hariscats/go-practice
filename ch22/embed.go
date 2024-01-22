package main

import "fmt"

type Logger struct {
	LogLevel string
}

func (l Logger) Log(message string) {
	fmt.Printf("[%s] %s\n", l.LogLevel, message)
}

type Server struct {
	Logger  // Embedding Logger into Server
	Address string
}

func main() {
	server := Server{
		Logger:  Logger{LogLevel: "INFO"},
		Address: "localhost:8080",
	}

	server.Log("Server started") // Method forwarding in action
}
