package task

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
