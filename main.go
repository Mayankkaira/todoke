package main

import "fmt"

func main() {
	for {
		fmt.Println("Welcome to Task Manager")
		fmt.Println("What is on your mind today")
		menuOptions := []string{
			"1. Add Task",
			"2. List Tasks",
			"3. Delete Task",
			"4. Check Status",
			"5. Exit",
		}
		for _, task := range menuOptions {
			fmt.Println(task)
		}
		fmt.Print("Choose an option:")
		var choice int
		fmt.Scan(&choice)
	
		switch choice {
		case 1:
			fmt.Println("Adding Task...")
		case 2:
			fmt.Println("Listing Tasks...")
		case 3:
			fmt.Println("Deleting Task...")
		case 4:
			fmt.Println("Cheking Status...")
		case 5:
			fmt.Println("Goodbye!")
			return 
			
		default:
			fmt.Println("Enter Valid Choice")
		}
	}

}
