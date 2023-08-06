package main

import (
	"log"
	"net/http"

	"github.com/BoskyWSMFN/go-lesson-multithreading-architecture/internal/architecture/mvc"
)

func main() {
	http.HandleFunc("/tasks", mvc.TasksHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
