package main

import (
	"bufio"
	"fmt"
	"os"
)

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
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(reader)
	for {
		for _, task := range menuOptions {
			fmt.Println(task)
		}
		fmt.Print("Choose an option:")
		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println(err)
			continue
		}
		switch choice {
		case 1:
			fmt.Println("Adding Task...")
		case 2:
			fmt.Println("Listing Tasks...")
		case 3:
			fmt.Println("Deleting Task...")
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
