package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Mayankkaira/todoke/storage"
	"github.com/Mayankkaira/todoke/task"
)

func main() {
	tasks, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Failed to load tasks:", err)
		fmt.Println("Exiting...")
		return
	}
	fmt.Println("Welcome to Task Manager")
	fmt.Println("What is on your mind today")
	menuOptions := []string{
		"1. Add Task",
		"2. List Tasks",
		"3. Delete Task",
		"4. Toggle Task Status",
		"5. Exit",
	}
	
	reader := bufio.NewReader(os.Stdin)
	for {
		for _, task := range menuOptions {
			fmt.Println(task)
		}
		fmt.Print("Choose an option:")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch choice {
		case 1:
			fmt.Println("enter the task name:")
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				continue
			}
			input = strings.TrimSpace(input)
			tasks, err = task.AddTask(tasks, input)
			if err != nil {
				fmt.Println("Failed to add task:", err)
				continue
			}
			err = storage.SaveTasks(tasks)
			if err != nil {
				fmt.Println("Failed to save tasks:", err, "\nYour changes are still available in this session.Fix the issue and try again before exiting.")
			} else {
				fmt.Println("Task added successfully.")
			}
		case 2:
			if len(tasks) == 0 {
				fmt.Println("No task found.")
				continue
			}
			printTasks(tasks)
		case 3:
			if len(tasks) == 0 {
				fmt.Println("There is no task to delete.")
				continue
			}
			printTasks(tasks)
			fmt.Println("enter the task no. u want to delete.")
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				continue
			}
			input = strings.TrimSpace(input)
			taskNumber, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if taskNumber < 1 || taskNumber > len(tasks) {
				fmt.Println("Invalid task number.")
				continue
			}
			index := taskNumber - 1
			tasks, err = task.DeleteTask(tasks, index)
			if err != nil {
				fmt.Println("Failed to delete task:", err)
				continue
			}
			err = storage.SaveTasks(tasks)
			if err != nil {
				fmt.Println("Failed to delete tasks:", err, "\nYour changes are still available in this session.Fix the issue and try again before exiting.")
			} else {
				fmt.Println("Task deleted successfully.")
			}
		case 4:
			if len(tasks) == 0 {
				fmt.Println("Task list is empty")
				continue
			}
			printTasks(tasks)
			fmt.Print("enter task no. to mark complete/incomplete:")
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				continue
			}
			input = strings.TrimSpace(input)
			taskNumber, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println(err)
				continue
			}
			if taskNumber < 1 || taskNumber > len(tasks) {
				fmt.Println("Invalid task number.")
				continue
			}
			index := taskNumber - 1
			err = task.ToggleTask(tasks, index)
			if err != nil {
				fmt.Println("Failed to toggle task:", err)
				continue
			}
			err = storage.SaveTasks(tasks)
			if err != nil {
				fmt.Println("Failed to save tasks:", err, "\nYour changes are still available in this session.Fix the issue and try again before exiting.")
			} else {
				fmt.Println("Task status updated.")
			}
		case 5:
			err := storage.SaveTasks(tasks)
			if err != nil {
				fmt.Println("Failed to save tasks:", err, "\nYour latest changes could not be saved and will be lost after exiting.")
			}
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid Option.")
		}

	}
}
