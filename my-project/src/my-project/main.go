package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}
type Node struct {
	task string
	Next *Node
}

type LinkedListNode struct {
	head   *Node
	tail   *Node
	length int
}

type Tree struct {
	Task       string
	ParentNode *Tree
	ChildNodes []*Tree
}

type Stack struct {
	trashedItems []string
}

var stack Stack

type Saved struct {
	savedItems []string
}

type List []item

type TaskManager struct {
	currentNode *Tree
	root        *Tree
	tasks       List
}

func (tm *TaskManager) Initialize() {
	tm.root = &Tree{
		Task:       "Root",
		ParentNode: nil,
		ChildNodes: []*Tree{},
	}
	tm.currentNode = tm.root
}

func (tm *TaskManager) AddTask(task string, LinkedList *LinkedListNode) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	tm.tasks = append(tm.tasks, t)

	node := &Node{
		task: task,
	}
	if LinkedList.tail == nil {
		LinkedList.head = node
		LinkedList.tail = node
	} else {
		LinkedList.tail.Next = node
		LinkedList.tail = node

	}
	LinkedList.length++
	printLinkedListThatWasGenerated(LinkedList)

	newTreeNode := &Tree{
		Task:       task,
		ParentNode: tm.currentNode,
		ChildNodes: []*Tree{},
	}
	tm.currentNode.ChildNodes = append(tm.currentNode.ChildNodes, newTreeNode)
	tm.currentNode = newTreeNode

	fmt.Println("Task added:", task)
}
func printLinkedListThatWasGenerated(LinkedList *LinkedListNode) {

	currentNode := LinkedList.head
	for currentNode != nil {
		fmt.Println(currentNode.task, "Print the LinkedList")
		currentNode = currentNode.Next

	}

}

func (tm *TaskManager) Undo() {
	if tm.currentNode.ParentNode != nil {
		tm.currentNode = tm.currentNode.ParentNode
		fmt.Println("Undo: Moved back to parent task:", tm.currentNode.Task)
		tm.PrintTreeFromRoot(tm.currentNode)
	} else {
		fmt.Println("Undo: Already at the root task.")
	}
}

func (tm *TaskManager) PrintTreeFromRoot(currentNode *Tree) {
	var path []string
	for n := currentNode; n != nil; n = n.ParentNode {
		path = append([]string{n.Task}, path...)
	}
	fmt.Println("Current task path:")
	for _, task := range path {
		fmt.Println("-", task)
	}
}

func (tm *TaskManager) Redo() {
	if len(tm.currentNode.ChildNodes) > 0 {
		tm.currentNode = tm.currentNode.ChildNodes[0]

		fmt.Println("Redo: Moved to the first child task:", tm.currentNode.Task)
		tm.PrintTreeFromRoot(tm.currentNode)
	} else {
		fmt.Println("Redo: No child tasks available.")
	}

}

func (tm *TaskManager) PrintTree(node *Tree, depth int) {
	// Print the entire tree, starting from the provided node
	fmt.Printf("%s- %s\n", strings.Repeat(" ", depth*2), node.Task)
	for _, child := range node.ChildNodes {
		tm.PrintTree(child, depth+1)
	}
}

func CompletedTask(l *List, i int, saved *Saved) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item does not exist")
	}
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()
	SaveTask(&ls[i-1], saved)
	fmt.Println("Task completed:", ls[i-1].Task)
	return nil
}

func SaveTask(task *item, saved *Saved) {
	saved.savedItems = append(saved.savedItems, task.Task)
	fmt.Println("Task saved:", task.Task)
}

func DeleteTask(saved *Saved) {
	TrashItems(saved, &stack)
	saved.savedItems = nil
	fmt.Println("The saved task(s) were deleted")
}

func TrashItems(saved *Saved, s *Stack) {
	for _, task := range saved.savedItems {
		s.trashedItems = append(s.trashedItems, task)
	}
	fmt.Println("Tasks moved to trash:", s.trashedItems)
}

func main() {
	tm := &TaskManager{}
	tm.Initialize()
	linkedList := &LinkedListNode{}
	saved := Saved{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Do you want to add a task? (yes/no):")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	if input == "yes" {
		fmt.Println("Enter your task:")
		task, _ := reader.ReadString('\n')
		task = strings.TrimSpace(task)
		tm.AddTask(task, linkedList)
	}

	for {
		fmt.Println("Enter command (add/complete/delete/list/trash/undo/redo/exit):")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(strings.ToLower(command))

		switch command {
		case "add":
			fmt.Println("Enter your task:")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)
			tm.AddTask(task, linkedList)
		case "complete":
			fmt.Println("Enter task index to complete:")
			indexStr, _ := reader.ReadString('\n')
			indexStr = strings.TrimSpace(indexStr)
			index, err := strconv.Atoi(indexStr)
			if err != nil {
				fmt.Println("Invalid index.")
				continue
			}
			if err := CompletedTask(&tm.tasks, index, &saved); err != nil {
				fmt.Println("Error:", err)
			}
		case "delete":
			DeleteTask(&saved)
		case "list":
			fmt.Println("Tasks:")
			for i, task := range tm.tasks {
				fmt.Printf("%d. %s (Done: %t)\n", i+1, task.Task, task.Done)
			}
		case "trash":
			fmt.Println("Trashed tasks:")
			for _, task := range stack.trashedItems {
				fmt.Println("-", task)
			}
		case "undo":
			tm.Undo()
		case "redo":
			tm.Redo()
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Unknown command. Available commands: add, complete, delete, list, trash, undo, redo, exit")
		}
	}
}
