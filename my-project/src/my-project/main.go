package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type item struct {
	Task        string
	Done        bool
	Duration    int
	Priority    int
	CreatedAt   time.Time
	CompletedAt time.Time
}
type Node struct {
	task string
	Next *Node
}

type LinkedListNode struct {
	head    *Node
	tail    *Node
	NodeMap map[string]*Node
	length  int
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

// i am gonna write a function that will priotrize task
// based on the priority and duration
// and somehow it will suggest me which task
// i should complete fast
// func PrioritySuggetions(list *List) {
// 	slices.Sort(list, func(i, j int) bool {
// 		if list[i].Priority == list[j].Priority {
// 			return list[i].Duration < list[j].Duration
// 		}
// 		return list[i].Priority > list[j].Priority
// 	})

// }
func PrioritySuggestions(list List) {
	slices.Sort(*list, func(a, b item) int {
		if a.Priority == b.Priority {
			return cmp.Compare(a.Duration, b.Duration) // Sort by duration if priorities are the same
		}
		return cmp.Compare(b.Priority, a.Priority) // Higher priority first
	})
}

// i have to find the task in o(1) Time
func FindTask(searchString string, Linklist *LinkedListNode) {

	// i am gonna use a hashmap for instant lookup
	result, notfound := Linklist.NodeMap[searchString]
	if !notfound {

		log.Fatalf("the Searched Task dosent exists")

	}
	log.Printf("The task` is Found", result.task)

}

func (tm *TaskManager) AddTask(task string, LinkedList *LinkedListNode, Duration int, priority int) {

	if LinkedList.NodeMap == nil { // ✅ Prevent panic
		LinkedList.NodeMap = make(map[string]*Node)
	}

	t := item{
		Task:        task,
		Done:        false,
		Duration:    Duration,
		Priority:    priority,
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

	//here  i am  using a HashMap to Track the Nodes
	LinkedList.NodeMap[task] = node
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

func MoveDown(ListNode *LinkedListNode, TargetTask string) {

	if ListNode.head == nil || ListNode.head.Next == nil {
		fmt.Println("Debug: List is empty or has only one node. No movement possible.")
		return
	}

	fmt.Println("Debug: Starting MoveDown for task:", TargetTask)

	fmt.Println("Debug: Initial linked list:")
	// I am finding the targeted Node In O(1) Time here
	targetNode, exisits := ListNode.NodeMap[TargetTask]

	if !exisits || targetNode.Next == nil {
		fmt.Println("the targeted Node coulnt bne found")
		return
	}

	first := targetNode
	second := targetNode.Next

	first.Next = second.Next
	second.Next = first
	ListNode.NodeMap[first.task] = first
	ListNode.NodeMap[second.task] = second

	if ListNode.head == first {
		ListNode.head = second
	}
	if first.Next == nil {
		ListNode.tail = first
	}

	// i am finiding the TargetNode In O(N) time ->Below code
	// for prev.Next != nil && prev.Next.Next != nil {
	// 	fmt.Println("Debug: Checking node:", prev.Next.task)
	// 	if prev.Next.task == TargetTask {
	// 		fmt.Println("Debug: Found target task:", prev.Next.task)
	// 		break
	// 	}
	// 	prev = prev.Next
	// }

	// if prev.Next == nil || prev.Next.Next == nil {
	// 	fmt.Println("Debug: Target task not found or can't be swapped.")
	// 	return
	// }

	// first := prev.Next
	// second := first.Next

	// fmt.Printf("Debug: Swapping nodes '%s' and '%s'\n", first.task, second.task)
	// prev.Next = second
	// first.Next = second.Next
	// second.Next = first

	// if ListNode.head == first {
	// 	ListNode.head = second
	// }
	// if first.Next == nil {
	// 	ListNode.tail = first
	// }

	// fmt.Println("Debug: Linked list after swapping:")
	printLinkedListThatWasGenerated(ListNode)
}
func MoveUp(ListNode *LinkedListNode, TargetTask string) {

	if ListNode.head == nil || ListNode.head.Next == nil {

		fmt.Println("the target  node is head or tail")
		return

	}

	dummy := &Node{Next: ListNode.head}
	prev := dummy

	for prev.Next != nil && prev.Next.Next != nil {

		if prev.Next.task == TargetTask {

			break

		}
		prev = prev.Next

	}
	if prev.Next == nil || prev.Next.Next == nil {
		return
	}

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
		fmt.Println("Enter command (add/complete/delete/list/trash/undo/redo/Set low Priority/exit):")
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
		case "set_low_priority":
			fmt.Println("Enter the task to move down:")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)

			MoveDown(linkedList, task)

		case "exit":
			fmt.Println("Exiting...")

		case "searchtask":

			searchedTask, _ := reader.ReadString('\n')
			searchedTask = strings.TrimSpace(searchedTask)
			FindTask(searchedTask, linkedList)

			fmt.Println("")
			return
		default:
			fmt.Println("Unknown command. Available commands: add, complete, delete, list, trash, undo, redo, Set low Priority exit")
		}
	}
}
