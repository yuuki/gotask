package task

import (
	"testing"
)

func TestNewTasks(t *testing.T) {
	tasks := NewTasks("y_uuki")
	if tasks == nil {
		t.Error("cannot new tasks")
	}

	if tasks.Owner != "y_uuki" {
		t.Error("invalid task name")
	}
}

func TestAdd(t *testing.T) {
	tasks := NewTasks("y_uuki")
	tasks.Add(NewTask(1, "gokyoto"))
	tasks.Add(NewTask(2, "gototakakura"))

	if len(tasks.Tasks) != 2 {
		t.Error("too few the number of tasks")
	}
}

func TestDelete(t *testing.T) {
	tasks := NewTasks("y_uuki")
	tasks.Add(NewTask(1, "gokyoto"))
	tasks.Add(NewTask(2, "gototakakura"))
	tasks.Add(NewTask(3, "gotosugari"))

	tasks.Delete(1)

	if len(tasks.Tasks) != 2 {
		t.Error("too few the number of tasks")
	}

	for _, task := range tasks.Tasks {
		if task.Id == 1 {
			t.Error("failed to delete")
		}
	}
}
