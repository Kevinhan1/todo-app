package main

import "fmt"

//Struct task
type Task struct {
	ID          int
	Title       string
	Description string
	Deadline    string
	Completed   bool
}

//Slice task
var tasks []Task
var nextID int = 1

//Function to add a task
func addTask(title, description, deadline string) {
	task := Task{
		ID:          nextID,
		Title:       title,
		Description: description,
		Deadline:    deadline,
		Completed:   false,
	}
	tasks = append(tasks, task)
	nextID++
	fmt.Println("Task added successfully!")
}

//List all tasks
func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	// Header tabel
	fmt.Printf("%-4s | %-15s | %-25s | %-12s | %-10s\n", "ID", "Title", "Description", "Deadline", "Status")
	fmt.Println("-------------------------------------------------------------------------------")

	// Isi tabel
	for _, task := range tasks {
		status := "Pending"
		if task.Completed {
			status = "✅ Done"
		}
		fmt.Printf("%-4d | %-15s | %-25s | %-12s | %-10s\n",
			task.ID, task.Title, task.Description, task.Deadline, status)
	}
}

// Function to mark a task as completed
func completeTask(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			return
		}
	}
	fmt.Println("Task not found.")
}

//find task by ID
func findTaskByID(id int) *Task {
	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i]
		}
	}
	fmt.Println("Task not found.")
	return nil
}

//Function to delete a task
func deleteTask(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Task deleted successfully!")
			return
		}
	}
	fmt.Println("Task not found.")
}
