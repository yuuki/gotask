package main

import (
	. "github.com/y-uuki/gotask/task"
	. "github.com/y-uuki/gotask/web"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	log.Printf("Starting server...")

	tasks := NewTasks("y_uuki")
	tasks.Add(NewTask(1, "buy the milk"))
	tasks.Add(NewTask(2, "east sushi"))
	tasks.Add(NewTask(3, "goto takakura2jyo"))
	tasks.Delete(1)

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/tasks", TasksHandler(tasks))
	http.HandleFunc("/api/tasks.json", TasksJSONHandler(tasks))
	http.ListenAndServe(":5000", nil)
}
