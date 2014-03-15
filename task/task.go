package task

import (
	"fmt"
)

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

func (task Task) String() string {
	return fmt.Sprintf("%d) %s (%t)", task.Id, task.Detail, task.Done)
}

func (task Task) Display() {
	fmt.Printf("%s\n", task)
	return
}

func (task *Task) Edit(detail string) {
	task.Detail = detail
}

func (task *Task) Finish() {
	task.Done = true
}
