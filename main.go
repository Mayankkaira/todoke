package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mayankkaira/taskmanager/storage"
	"github.com/mayankkaira/taskmanager/task"
)

// type Task struct {
// 	Name      string
// 	Completed bool
// }

// func SaveTasks(tasks []Task) error {
// 	var content strings.Builder
// 	file, err := os.Create("tasks.txt")
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()
// 	for _, value := range tasks {
// 		content.WriteString(value.Name)
// 		content.WriteString("|")
// 		content.WriteString(strconv.FormatBool(value.Completed))
// 		content.WriteString("\n")
// 	}
// 	_, err = file.WriteString(content.String())
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func LoadTasks() ([]Task, error) {
// 	file, err := os.Open("tasks.txt")
// 	if err != nil {
// 		if os.IsNotExist(err) {
// 			return []Task{}, nil
// 		}
// 		return nil, err
// 	}
// 	defer file.Close()
// 	tasks := []Task{}
// 	reader := bufio.NewReader(file)
// 	for {
// 		data, err := reader.ReadString('\n')
// 		if data != "" {
// 			data = strings.TrimSpace(data)
// 			parts := strings.Split(data, "|")
// 			if len(parts) != 2 {
// 				return nil, fmt.Errorf("invalid task format")
// 			}
// 			part, err := strconv.ParseBool(parts[1])
// 			if err != nil {
// 				return nil, err
// 			}
// 			task := Task{Name: parts[0], Completed: part}
// 			tasks = append(tasks, task)
// 		}

// 		if err == io.EOF {
// 			return tasks, nil
// 		}
// 		if err != nil {
// 			return nil, err
// 		}

// 	}
// }

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
	// var tasks []Task
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
			tasks=task.AddTasks(tasks, reader)
			err = storage.SaveTasks(tasks)
			if err != nil {
				fmt.Println("Failed to save tasks:", err, "\nYour changes are still available in this session.Fix the issue and try again before exiting.")
			} else {
				fmt.Println("Task added successfully.")
			}
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
			tasks,err=task.DeleteTask(tasks,reader)
			if err!=nil{
				fmt.Println(err)
				continue
			}
			err=storage.SaveTasks(tasks)
			if err != nil {
				fmt.Println("Failed to delete tasks:", err, "\nYour changes are still available in this session.Fix the issue and try again before exiting.")
			} else{
				fmt.Println("Task deleted Successfully.")
			}
			// if len(tasks) == 0 {
			// 	fmt.Println("There is no task to delete")
			// 	continue
			// }
			// for i, value := range tasks {
			// 	fmt.Printf("%d. %s\n", i+1, value.Name)
			// }
			// fmt.Println("enter the task no. u want to delete")
			// input, err := reader.ReadString('\n')
			// if err != nil {
			// 	fmt.Println(err)
			// 	continue
			// }
			// input = strings.TrimSpace(input)
			// taskNumber, err := strconv.Atoi(input)
			// if err != nil {
			// 	fmt.Println(err)
			// 	continue
			// }
			// // if 1 <= taskNumber && taskNumber <= len(tasks) {
			// // 	index := taskNumber - 1
			// // 	left := tasks[:index]
			// // 	right := tasks[index+1:]
			// // 	tasks = append(left, right...)
			// // 	// fmt.Println(tasks)
			// // 	fmt.Println("Task deleted successfully")
			// // } else {
			// // 	fmt.Println("Invalid task number.")
			// // }
			// if taskNumber < 1 || taskNumber > len(tasks) {
			// 	fmt.Println("Invalid task number.")
			// 	continue
			// }
			// index := taskNumber - 1
			// left := tasks[:index]
			// right := tasks[index+1:]
			// tasks = append(left, right...)
			// err = storage.SaveTasks(tasks)
			// if err != nil {
			// 	fmt.Println("Failed to delete tasks:", err, "\nYour changes are still available in this session.Fix the issue and try again before exiting.")
			// } else {
			// 	fmt.Println("Task deleted Successfully.")
			// }
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
			err = storage.SaveTasks(tasks)
			if err != nil {
				fmt.Println("Failed to save tasks:", err, "\nYour changes are still available in this session.Fix the issue and try again before exiting.")
			} else {
				fmt.Println("Task status updated.")
			}
			//fmt.Println(index)
			// if !tasks[index].Completed {
			// 	tasks[index].Completed = true
			// 	fmt.Println("Task marked")
			// } else {
			// 	tasks[index].Completed = false
			// 	fmt.Println("Task marked as pending")
			// }
		case 5:
			err := storage.SaveTasks(tasks)
			if err != nil {
				fmt.Println("Failded to save tasks:", err, "\nYour latest changes could not be saved and will be lost after exiting.")
			}
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid Option")
		}

	}
}
