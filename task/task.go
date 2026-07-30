package task

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Task struct {
	Name      string
	Completed bool
}

func AddTasks(tasks []Task, r *bufio.Reader) []Task {
	fmt.Println("enter the task name u want to enter")
	input, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		//continue
		// return err
	}
	input = strings.TrimSpace(input)
	newTask := Task{Name: input, Completed: false}
	tasks = append(tasks, newTask)
	// err = storage.SaveTasks(tasks)
	// if err != nil {
	// 	fmt.Errorf("Failed to save tasks:", err, "\nYour changes are still available in this session.Fix the issue and try again before exiting.")
	// 	// fmt.Println("Your changes are still available in this session.Fix the issue and try again before exiting.")
	// }
	//return nil
	return tasks
}

func DeleteTask(tasks []Task, r *bufio.Reader) ([]Task, error) {
	if len(tasks) == 0 {
		// fmt.Println("There is no task to delete")
		return nil, fmt.Errorf("There is no task to delete")
	}
	for i, value := range tasks {
		fmt.Printf("%d. %s\n", i+1, value.Name)
	}
	fmt.Println("enter the task no. u want to delete")
	input, err := r.ReadString('\n')
	if err != nil {
		// fmt.Println(err)
		// continue
		return nil, err
	}
	input = strings.TrimSpace(input)
	taskNumber, err := strconv.Atoi(input)
	if err != nil {
		// fmt.Println(err)
		return nil, err
	}
	// if 1 <= taskNumber && taskNumber <= len(tasks) {
	// 	index := taskNumber - 1
	// 	left := tasks[:index]
	// 	right := tasks[index+1:]
	// 	tasks = append(left, right...)
	// 	// fmt.Println(tasks)
	// 	fmt.Println("Task deleted successfully")
	// } else {
	// 	fmt.Println("Invalid task number.")
	// }
	if taskNumber < 1 || taskNumber > len(tasks) {
		// fmt.Println("Invalid task number.")
		return nil, fmt.Errorf("Invalid task number")
	}
	index := taskNumber - 1
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
