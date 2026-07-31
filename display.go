package main

import (
	"fmt"

	"github.com/mayankkaira/taskmanager/task"
)

func printTasks(tasks []task.Task) {
	fmt.Println("Your Tasks")
	completeLabel := "✔"
	pendingLabel := " "
	for i, value := range tasks {
		// fmt.Printf("%T\n",tasks[i])
		if value.Completed {
			fmt.Printf("%d. [%s] %s\n", i+1, completeLabel, value.Name)
		} else {
			fmt.Printf("%d. [%s] %s\n", i+1, pendingLabel, value.Name)
		}
	}
}
