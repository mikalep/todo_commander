package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Todos []Todo

type Todo struct {
	Todo      string `json:"todo"`
	Completed bool   `json:"completed"`
}

func listOptions() {

	options := []string{"Show all todos", "Toggle todo status", "Add a new todo", "Edit todo", "Delete todo", "Exit"}

	for k, v := range options {
		fmt.Println(k+1, "=>", v)
	}

}

func showAllTodos() {

	fmt.Println("Here are your current todos:")

	todos := readTodos()

	json, err := json.MarshalIndent(todos, "", "    ")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json))

}

func toggleTodoStatus() {

	var input int

	allTodos := readTodos()

	fmt.Println("Which todo entry would you like toggle?\n")

	for i, v := range allTodos {
		fmt.Println(i+1, v.Todo)
	}

	fmt.Scanln(&input)

	allTodos[input-1].Completed = !allTodos[input-1].Completed

	WriteTodos(allTodos)

	fmt.Printf("Entry number %d marked as complete, well done!\n", input)

	showAllTodos()

}

func addNewTodo() {

	var newTodo Todo

	allTodos := readTodos()

	fmt.Println("Name of todo:")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSuffix(input, "\n")

	newTodo.Todo = input

	allTodos = append(allTodos, newTodo)

	WriteTodos(allTodos)

}

func editTodo() {
	fmt.Println("Edit todo")
}

func deleteTodo() {

	var input int

	allTodos := readTodos()

	for i, v := range allTodos {
		fmt.Println(i+1, v.Todo)
	}

	fmt.Println("Which todo would you like to remove?\n")

	fmt.Scanln(&input)

	allTodos = append(allTodos[:input-1], allTodos[input:]...)

	fmt.Printf("Entry number %d was successfully deleted.\n", input)

	WriteTodos(allTodos)

}

func WriteTodos(allTodos Todos) {

	b, err := json.MarshalIndent(allTodos, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("todo.json", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func chooseAction() {

	run := true

	for run {

		listOptions()

		var choice int
		fmt.Scanln(&choice)

		switch {
		case choice == 1:
			showAllTodos()
		case choice == 2:
			toggleTodoStatus()
		case choice == 3:
			addNewTodo()
		case choice == 4:
			editTodo()
		case choice == 5:
			deleteTodo()
		case choice == 6:
			run = false
			fmt.Println("Bye friend!")
		default:
			fmt.Println("No such option! Please try again.")
		}
	}

}

func readTodos() Todos {

	jsonFile, err := os.Open("todo.json")

	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	openedJsonFile, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Fatal(err)
	}

	var todos Todos

	err = json.Unmarshal(openedJsonFile, &todos)

	if err != nil {
		log.Fatal(err)
	}

	return todos
}

func main() {

	fmt.Println("Hey there friend, welcome to To-Do Commander, press a number to perform the desired action: \n")

	chooseAction()

}
