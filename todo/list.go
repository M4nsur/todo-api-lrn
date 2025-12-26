package todo

import "fmt"

type List struct {
	tasks map[string]Task
}

func NewList() *List {
	return &List{
		tasks: make(map[string]Task),
	}
}


func (l *List) AddTask(task Task) error {
	if _, ok := l.tasks[task.Title]; ok {
		return ErrTaskAlreadyExists
	}
	l.tasks[task.Title] = task
	return nil
}

func (l *List) ListTasks() map[string]Task {
	return l.tasks
}

func (l *List) DoneTask(title string) {
	task, ok := l.tasks[title]
	if(!ok) {
		fmt.Println(ErrTaskNotFound)

	}

	task.Done()

	l.tasks[title] = task
}


func (l *List) DeleteTask(title string) string {
	_, ok := l.tasks[title]

	if(!ok) {
		fmt.Println(ErrTaskNotFound)
	}

	delete(l.tasks, title)

	return ""
}