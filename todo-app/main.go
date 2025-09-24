package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {
		fmt.Println("=== To-Do List Application Menu ===")
		fmt.Println("1. Add Task")
		fmt.Println("2. List	Tasks")
		fmt.Println("3. Complete Task")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")

		choice := input("Choose an option: ")

		switch choice {
		case "1":
			title := input("Enter task title: ")
			description := input("Enter task description: ")
			deadline := input("Enter task deadline: ")
			addTask(title, description, deadline)
		case "2":
			listTasks()
		case "3":
			idStr := input("Enter task ID to complete: ")
			id, err := strconv.Atoi(idStr)
			if err == nil {
				completeTask(id)
				fmt.Println("Task	marked as completed!")
			} else {
				fmt.Println("Invalid task ID ")
			}

		case "4":
			idStr := input("Enter task ID to delete: ")
			id, err := strconv.Atoi(idStr)
			if err == nil {
				deleteTask(id)
				fmt.Println("Task deleted successfully!")
			} else {
				fmt.Println("Invalid task ID")
			}

		case "5":
			fmt.Println("Exiting application. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
