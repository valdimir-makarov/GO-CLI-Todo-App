package main

import "fmt"

func main() {
	var tasks List
	var saved Saved

	AddTask(&tasks, "Buy groceries")
	AddTask(&tasks, "Learn Go")
	AddTask(&tasks, "Write code")

	fmt.Println("Tasks after adding:")
	for i, task := range tasks {
		fmt.Printf("%d. %s (Done: %t)\n", i+1, task.Task, task.Done)
	}

	if err := CompletedTask(&tasks, 1, &saved); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\nTasks after completing one:")
	for i, task := range tasks {
		fmt.Printf("%d. %s (Done: %t)\n", i+1, task.Task, task.Done)

	}
	DeleteTask(&saved)
}
