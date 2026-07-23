package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Name      string
	Completed bool
}

func SaveTasks(tasks []Task) error {
	var content strings.Builder
	file, err := os.Create("tasks.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	for _, value := range tasks {
		content.WriteString(value.Name)
		content.WriteString(",")
		content.WriteString(strconv.FormatBool(value.Completed))
		content.WriteString("\n")
	}
	count, err := file.WriteString(content.String())
	if err != nil {
		return err
	}
	_=count
	// fmt.Println(Count)
	return  nil
}
func main() {
	fmt.Println("Welcome to Task Manager")
	fmt.Println("What is on your mind today")
	menuOptions := []string{
		"1. Add Task",
		"2. List Tasks",
		"3. Delete Task",
		"4. Toggle Task Status",
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
			newTask := Task{Name: input, Completed: false}
			tasks = append(tasks, newTask)
			fmt.Println("Task added Successfully...")
			// fmt.Println(tasks)

		case 2:
			if len(tasks) == 0 {
				fmt.Println("No task found")
				continue
			}
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
		case 3:
			if len(tasks) == 0 {
				fmt.Println("There is no task to delete")
				continue
			}
			for i, value := range tasks {
				fmt.Printf("%d. %s\n", i+1, value.Name)
			}
			fmt.Println("enter the task no. u want to delete")
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
				fmt.Println("Invalid task number.")
				continue
			}
			index := taskNumber - 1
			left := tasks[:index]
			right := tasks[index+1:]
			tasks = append(left, right...)
			fmt.Println("Task deleted successfully...")
		case 4:
			// fmt.Println("Checking Status...")
			if len(tasks) == 0 {
				fmt.Println("Task list is empty")
				continue
			}
			fmt.Println("Task list")
			completeLabel := "✔"
			pendingLabel := " "
			for i, value := range tasks {
				// fmt.Printf("%d.%s[%v]\n", i+1, value.Name, value.Complete)
				if value.Completed {
					fmt.Printf("%d. [%s] %s\n", i+1, completeLabel, value.Name)
				} else {
					fmt.Printf("%d. [%s] %s\n", i+1, pendingLabel, value.Name)
				}
			}
			fmt.Print("enter task no.")
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
			tasks[index].Completed = !tasks[index].Completed
			fmt.Println("Task status updated...")
			//fmt.Println(index)
			// if !tasks[index].Completed {
			// 	tasks[index].Completed = true
			// 	fmt.Println("Task marked")
			// } else {
			// 	tasks[index].Completed = false
			// 	fmt.Println("Task marked as pending")
			// }

		case 5:
			err := SaveTasks(tasks)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid Option")
		}

	}

}
