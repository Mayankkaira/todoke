package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Name     string
	Complete bool
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
			newTask := Task{Name: input}
			tasks = append(tasks, newTask)
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
				fmt.Printf("%d.%s [%v]\n", i+1, value.Name,value.Complete)
			}
		case 3:
			if len(tasks) == 0 {
				fmt.Println("There is no task to delete")
				continue
			}
			for i, value := range tasks {
				fmt.Printf("%d.%s\n", i+1, value.Name)
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
			fmt.Println("Task deleted successfully..")
		case 4:
			// fmt.Println("Checking Status...")
			if len(tasks)==0{
				fmt.Println("Task list is empty")
				continue
			}
			fmt.println("Task list")
			for i,value:=range tasks{
				fmt.Printf("%d.%s[%v]",i+1,value.Name,value.Complete)
			}
			input,err:=reader.ReadString('\n')
			if err!=nil{
				fmt.Println(err)
				continue
			}
			input=strings.TrimSpace(input)
			taskNumber,err:=strconv.Atoi(input)
			if err!=nil{
				fmt.Println(err)
				continue
			}
			if taskNumber<1 || taskNumber> len(tasks){
				fmt.Println("Invalid task number.")
			}
			index:=taskNumber-1
			fmt.Println(index)
		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid Option")
		}

	}

}
