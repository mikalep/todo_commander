package main

import (
	"fmt"
)

// List all available options for the user
func listOptions() {

	options := []string{"Show all tasks", "Change task status", "Add a new task", "Edit task", "Delete task", "Exit"}

	for k, v := range options {
		fmt.Println(k+1, "=>", v)
	}

}

// Functions to handle each specific task
func showAllTasks() {
	fmt.Println("All tasks")
}

func changeTaskStatus() {
	fmt.Println("Status changed.")
}

func addNewTask() {
	fmt.Println("New task")
}

func editTask() {
	fmt.Println("Edit task")
}

func deleteTask() {
	fmt.Println("Delete task")
}

// Handle the actual user choice
func chooseAction() {

	listOptions()

	run := true

	for run {

		var choice int
		fmt.Scanln(&choice)

		switch {
		case choice == 1:
			showAllTasks()
		case choice == 2:
			changeTaskStatus()
		case choice == 3:
			addNewTask()
		case choice == 4:
			editTask()
		case choice == 5:
			deleteTask()
		case choice == 6:
			run = false
		default:
			fmt.Println("No such option! Please try again.")
			listOptions()
		}
	}

}

func main() {

	fmt.Println("Hey there friend, welcome to To-Do Commander, press a number to perform the desired action: \n")

	chooseAction()

}
