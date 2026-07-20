package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

)

type Task struct {
	Name string
}

func main() {
	fmt.Println("Welcome to Task Manager")
	fmt.Println("What is on your mind today")
	menuOptions := []string{
		"1. Add Task",
		"2. List Tasks",
		"3. Delete Task",
		"4. Check Status",
		"5. Exit",
	}
	var tasks []Task
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
		// fmt.Printf("%q\n", input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch choice {
		case 1:
			fmt.Println("enter the task name u want to enter")
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				continue
			}

			input = strings.TrimSpace(input)
			a := Task{Name: input}
			tasks = append(tasks, a)
			fmt.Println("Task added Successfully...")
			// fmt.Println(tasks)

		case 2:
			if len(tasks) == 0 {
				fmt.Println("No task found")
				continue
			}
			fmt.Println("Your Tasks")
			for i, value := range tasks {
				// fmt.Printf("%T\n",tasks[i])
				fmt.Printf("%d.%s\n", i+1, value.Name)
			}
		case 3:
			if len(tasks)==0{
				fmt.Println("There is no task to delete")
				continue
			}
			for i,value:=range tasks{
				fmt.Printf("%d.%s\n",i+1,value.Name)
			}
			fmt.Println("enter the task no u want to delete")
			input,err:=reader.ReadString('\n')
			if err!=nil{
				fmt.Println(err)
				continue
			}
			input=strings.TrimSpace(input)
			tasks=remove()

		case 4:
			fmt.Println("Checking Status...")
		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid Option")
		}

	}

}
