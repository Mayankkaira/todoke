package main

import (
	"fmt"

	"github.com/Mayankkaira/todoke/task"
)

func printTasks(tasks []task.Task) {
	fmt.Println("Your Tasks")
	const (
		completeLabel = "✔"
		pendingLabel  = " "
	)
	for i, value := range tasks {
		if value.Completed {
			fmt.Printf("%d. [%s] %s\n", i+1, completeLabel, value.Name)
		} else {
			fmt.Printf("%d. [%s] %s\n", i+1, pendingLabel, value.Name)
		}
	}
}
