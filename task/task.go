package task

import "fmt"

type Task struct {
	Name      string
	Completed bool
}

func AddTask(tasks []Task, name string) []Task {
	newTask := Task{Name: name, Completed: false}
	tasks = append(tasks, newTask)
	// err = storage.SaveTasks(tasks)
	// if err != nil {
	// 	fmt.Errorf("Failed to save tasks:", err, "\nYour changes are still available in this session.Fix the issue and try again before exiting.")
	// 	// fmt.Println("Your changes are still available in this session.Fix the issue and try again before exiting.")
	// }
	//return nil
	return tasks
}

func DeleteTask(tasks []Task, index int) ([]Task, error) {
	if index < 0 || index >= len(tasks) {
		return tasks, fmt.Errorf("Invalid task number")
	}
	left := tasks[:index]
	right := tasks[index+1:]
	tasks = append(left, right...)
	return tasks, nil
}

// err = storage.SaveTasks(tasks)
// if err != nil {
// 	fmt.Println("Failed to delete tasks:", err, "\nYour changes are still available in this session.Fix the issue and try again before exiting.")
// } else {
// 	fmt.Println("Task deleted Successfully.")
// }

func ToggleTask(tasks []Task, index int) {
	tasks[index].Completed = !tasks[index].Completed
}
