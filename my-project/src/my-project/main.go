package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var tasks List
	var saved Saved

	// Check if command-line arguments are provided
	if len(os.Args) > 1 {
		command := os.Args[1]

		switch command {
		case "add":
			if len(os.Args) < 3 {
				fmt.Println("Usage: add <task>")
				return
			}
			task := os.Args[2]
			AddTask(&tasks, task)
			fmt.Println("Task added:", task)
		case "complete":
			if len(os.Args) < 3 {
				fmt.Println("Usage: complete <task_index>")
				return
			}
			index, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println("Invalid task index:", os.Args[2])
				return
			}
			if err := CompletedTask(&tasks, index, &saved); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Task %d marked as completed.\n", index)
			}
		case "delete":
			DeleteTask(&saved)
		case "list":
			fmt.Println("Tasks:")
			for i, task := range tasks {
				fmt.Printf("%d. %s (Done: %t)\n", i+1, task.Task, task.Done)
			}
		case "trash":
			fmt.Println("Trashed tasks:")
			for _, task := range stack.trashedItems {
				fmt.Println("-", task)
			}
		default:
			fmt.Println("Unknown command. Available commands: add, complete, delete, list, trash")
		}
	} else {
		fmt.Println("No command provided. Usage: <command> [arguments]")
		fmt.Println("Available commands: add, complete, delete, list, trash")
	}
}
