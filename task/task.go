package task

import (
	"errors"
	"strings"
)

type Task struct {
	Name      string
	Completed bool
}

func AddTask(tasks []Task, name string) ([]Task, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return tasks, errors.New("task name cannot be empty")
	}
	newTask := Task{Name: name, Completed: false}
	tasks = append(tasks, newTask)
	return tasks, nil
}

func DeleteTask(tasks []Task, index int) ([]Task, error) {
	if index < 0 || index >= len(tasks) {
		return tasks, errors.New("task index out of range")
	}
	left := tasks[:index]
	right := tasks[index+1:]
	tasks = append(left, right...)
	return tasks, nil
}

func ToggleTask(tasks []Task, index int) error {
	if index < 0 || index >= len(tasks) {
		return errors.New("task index out of range")
	}
	tasks[index].Completed = !tasks[index].Completed
	return nil
}
