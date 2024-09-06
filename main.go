package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"todoList/todo"
)

const todosFile = "todos.json"

func main() {
	// Define subcommands
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addTask := addCmd.String("task", "", "The task to be added")

	if len(os.Args) < 2 {
		fmt.Println("Expected 'add', 'update', or 'list' subcommand.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		handleAdd(*addTask)
	case "delete":
		handleDelete(os.Args[2:])
	case "update", "mark-done", "mark-in-progress":
		handleUpdate(os.Args[1], os.Args[2:])
	case "list":
		handleList(os.Args[2:])
	default:
		fmt.Println("Unknown command:", os.Args[1])
		os.Exit(1)
	}
}

func handleAdd(task string) {
	if task == "" {
		fmt.Println("Please provide a task description.")
		os.Exit(1)
	}

	todos, err := todo.LoadTodos(todosFile)
	if err != nil {
		fmt.Println("Error loading todos:", err)
		os.Exit(1)
	}

	todos.AddTodo(task)

	if err := todos.SaveTodos(todosFile); err != nil {
		fmt.Println("Error saving todo:", err)
		os.Exit(1)
	}

	fmt.Println("Added new todo:", task)
}

func handleUpdate(command string, args []string) {
	if len(args) < 1 {
		fmt.Println("Please provide the ID of the todo to update.")
		os.Exit(1)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid ID. Please provide a numeric ID.")
		os.Exit(1)
	}

	task := ""
	status := ""
	if command == "mark-done" {
		status = todo.StatusDone
	} else if command == "mark-in-progress" {
		status = todo.StatusInProgress
	} else if len(args) > 1 {
		task = args[1]
	}

	todos, err := todo.LoadTodos(todosFile)
	if err != nil {
		fmt.Println("Error loading todos:", err)
		os.Exit(1)
	}

	if err := todos.UpdateTodo(id, task, status); err != nil {
		fmt.Println("Error updating todo:", err)
		os.Exit(1)
	}

	if err := todos.SaveTodos(todosFile); err != nil {
		fmt.Println("Error saving todo:", err)
		os.Exit(1)
	}

	fmt.Println("Updated todo:", id)
}

func handleDelete(args []string) {
	if len(args) < 1 {
		fmt.Println("Please provide the ID of the todo to delete.")
		os.Exit(1)
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid ID. Please provide a numeric ID.")
		os.Exit(1)
	}

	todos, err := todo.LoadTodos(todosFile)
	if err != nil {
		fmt.Println("Error loading todos:", err)
		os.Exit(1)
	}

	if err := todos.DeleteTodo(id); err != nil {
		fmt.Println("Error deleting todo:", err)
		os.Exit(1)
	}

	if err := todos.SaveTodos(todosFile); err != nil {
		fmt.Println("Error saving todo:", err)
		os.Exit(1)
	}

	fmt.Println("Deleted todo:", id)
}

func handleList(args []string) {
	status := ""
	if len(args) >= 1 {
		status = args[0]
	}

	todos, err := todo.LoadTodos(todosFile)
	if err != nil {
		fmt.Println("Error loading todos:", err)
		os.Exit(1)
	}

	todos.ListTodos(status)
}
