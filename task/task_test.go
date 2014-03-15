package task

import (
	"testing"
)

func TestNewTask(t *testing.T) {
	task := NewTask(1, "gokyoto")
	if task == nil {
		t.Error("cannot new task")
	}

	if task.Id != 1 {
		t.Error("invalid task id")
	}

	if task.Detail != "gokyoto" {
		t.Error("invalid task id")
	}
}

func TestEdit(t *testing.T) {
	task := NewTask(2, "gototakakura")

	task.Edit("gotosugari")

	if task.Detail != "gotosugari" {
		t.Error("failed to edit")
	}

	if task.Id != 2 {
		t.Error("should not modify id")
	}
}
