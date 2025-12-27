package todo

import (
	"maps"
)

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
	tmp := make(map[string]Task, len(l.tasks))

	maps.Copy(tmp, l.tasks)

	return tmp
}

func (l *List) CompleteTask(title string) error {
	task, ok := l.tasks[title]
	if(!ok) {
		return ErrTaskNotFound
	}

	task.Complete()

	l.tasks[title] = task
	return nil
}


func (l *List) DeleteTask(title string) error {
	_, ok := l.tasks[title]

	if(!ok) {
		return ErrTaskNotFound
	}

	delete(l.tasks, title)

	return nil
}

func (l *List) ListNotCompletedTasks () map[string]Task {
	notCompleted := make(map[string]Task)

	for k, v := range l.tasks {
		if(!v.Completed) {
			notCompleted[k] = v
		}
	}

	return notCompleted
}