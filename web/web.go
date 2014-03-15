package web

import (
	"encoding/json"
	"fmt"
	. "github.com/y-uuki/gotask/task"
	"html/template"
	"log"
	"net/http"
)

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
