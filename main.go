package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

type Task struct {
	Id     int    `json:"id"`
	Detail string `json:"details"`
	Done   bool   `json:"done"`
}

func NewTask(id int, detail string) *Task {
	task := &Task{
		Id:     id,
		Detail: detail,
	}
	return task
}

func (task Task) String() string { // 引数なし,戻り値あり
	return fmt.Sprintf("%d) %s (%t)", task.Id, task.Detail, task.Done)
}

func (task Task) Display() {
	fmt.Printf("%s\n", task)
	return
}

func (task *Task) Edit(detail string) {
	task.Detail = detail
}

func (task *Task) Finish() { // 引数,戻り値無し
	task.Done = true
}

type Tasks struct {
	Owner string  `json:"owner"`
	Tasks []*Task `json:"tasks"`
}

func NewTasks(owner string) *Tasks {
	tasks := &Tasks{
		Owner: owner,
	}
	return tasks
}

func (tasks *Tasks) Add(task *Task) {
	tasks.Tasks = append(tasks.Tasks, task)
}

func (tasks *Tasks) Delete(id int) {
	for i, task := range tasks.Tasks {
		if task.Id == id {
			copy(tasks.Tasks[i:], tasks.Tasks[i+1:])
			tasks.Tasks[len(tasks.Tasks)-1] = nil
			tasks.Tasks = tasks.Tasks[:len(tasks.Tasks)-1]
			break
		}
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// heredoc
	html := `
  <!DOCTYPE html>
  <title>tasks</title>
  <h1>here is your <a href="/tasks">tasks</a></h1>
  `
	fmt.Fprintf(w, html)
}

func TasksHandler(tasks *Tasks) func(w http.ResponseWriter, r *http.Request) {
	html := `
  <!DOCTYPE html>
    <title>tasks</title>
    <h1>{{.Owner}}'s tasks</h1>
    <ul>
        {{range .Tasks}}
        <li>{{.}}</li>
        {{end}}
    </ul>
    `

	// 第二引数がエラーを返すイディオム
	TasksList, err := template.New("tasklist").Parse(html)
	if err != nil {
		log.Fatal(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		TasksList.Execute(w, tasks)
	}
}

func TasksJSONHandler(tasks *Tasks) func(w http.ResponseWriter, r *http.Request) {
	b, err := json.MarshalIndent(tasks, "", "  ") // プレフィックス、インデントあり
	if err != nil {
		log.Fatal(err)
	}
	json := string(b)
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, json)
	}
}

func main() {
	log.Printf("Starting server...")

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/tasks", TasksHandler(tasks))
	http.HandleFunc("/api/tasks.json", TasksJSONHandler(tasks))
	http.ListenAndServe(":5000", nil)
}
