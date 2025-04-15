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

	"github.com/google/uuid"
)

type item struct {
	ID          uuid.UUID
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

type Graph struct {
	adjlist map[string][]string
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
	graph       *Graph
}

func (tm *TaskManager) Initialize() {
	tm.root = &Tree{
		Task:       "Root",
		ParentNode: nil,
		ChildNodes: []*Tree{},
	}
	tm.currentNode = tm.root
	tm.graph = NewGraph()
}

func NewGraph() *Graph {
	return &Graph{
		adjlist: make(map[string][]string),
	}
}

func PrioritySuggestions(list *List) {
	if len(*list) == 0 {
		log.Println("List is empty")
		return
	}
	slices.SortFunc(*list, func(a, b item) int {
		if a.Priority == b.Priority {
			return cmp.Compare(a.Duration, b.Duration)
		}
		return cmp.Compare(b.Priority, a.Priority)
	})

	fmt.Println("Suggested Work Priority List:")
	for _, it := range *list {
		fmt.Printf("Task: %s | Duration: %d | Priority: %d | Created At: %s\n",
			it.Task, it.Duration, it.Priority, it.CreatedAt.Format(time.RFC3339))
	}
}

func FindTask(searchString string, LinkedList *LinkedListNode) {
	if result, ok := LinkedList.NodeMap[searchString]; ok {
		log.Printf("Task Found: %s", result.task)
	} else {
		log.Println("Task not found")
	}
}

func (tm *TaskManager) AddTask(task string, LinkedList *LinkedListNode, duration, priority int, DependentTasks []string) {
	if LinkedList.NodeMap == nil {
		LinkedList.NodeMap = make(map[string]*Node)
	}

	t := item{
		ID:        uuid.New(),
		Task:      task,
		Duration:  duration,
		Priority:  priority,
		CreatedAt: time.Now(),
	}
	tm.tasks = append(tm.tasks, t)
	tm.graph.adjlist[t.Task] = append(tm.graph.adjlist[t.Task], DependentTasks...)

	node := &Node{task: task}
	if LinkedList.tail == nil {
		LinkedList.head = node
	} else {
		LinkedList.tail.Next = node
	}
	LinkedList.tail = node
	LinkedList.NodeMap[task] = node
	LinkedList.length++

	newTreeNode := &Tree{
		Task:       task,
		ParentNode: tm.currentNode,
	}
	tm.currentNode.ChildNodes = append(tm.currentNode.ChildNodes, newTreeNode)
	tm.currentNode = newTreeNode

	fmt.Println("Task added:", task)
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

func CompletedTask(l *List, i int, saved *Saved) error {
	if i <= 0 || i > len(*l) {
		return fmt.Errorf("item does not exist")
	}
	(*l)[i-1].Done = true
	(*l)[i-1].CompletedAt = time.Now()
	SaveTask(&(*l)[i-1], saved)
	fmt.Println("Task completed:", (*l)[i-1].Task)
	return nil
}

func SaveTask(task *item, saved *Saved) {
	saved.savedItems = append(saved.savedItems, task.Task)
	fmt.Println("Task saved:", task.Task)
}

func (tm *TaskManager) DeleteTask(saved *Saved, indexReal int) {
	TrashItems(saved, &stack)
	index := indexReal - 1
	if index > 0 && index < len(tm.tasks) {
		tm.tasks = append(tm.tasks[:index], tm.tasks[index+1:]...)

	}

	saved.savedItems = nil
	fmt.Println("The saved task(s) were deleted")
}

func TrashItems(saved *Saved, s *Stack) {
	s.trashedItems = append(s.trashedItems, saved.savedItems...)
	fmt.Println("Tasks moved to trash:", s.trashedItems)
}

func main() {
	tm := &TaskManager{}
	tm.Initialize()
	linkedList := &LinkedListNode{}
	saved := &Saved{}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter command (add/complete/delete/list/trash/undo/redo/suggestpriority/searchtask/exit):")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(strings.ToLower(command))

		switch command {
		case "add":
			fmt.Println("Enter your task:")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)

			fmt.Println("Enter dependent tasks (comma separated):")
			dependentTask, _ := reader.ReadString('\n')
			dependentTask = strings.TrimSpace(dependentTask)

			fmt.Print("Enter duration: ")
			durationStr, _ := reader.ReadString('\n')
			duration, err := strconv.Atoi(strings.TrimSpace(durationStr))
			if err != nil {
				fmt.Println("Invalid duration input")
				continue
			}

			fmt.Print("Enter priority: ")
			priorityStr, _ := reader.ReadString('\n')
			priority, err := strconv.Atoi(strings.TrimSpace(priorityStr))
			if err != nil {
				fmt.Println("Invalid priority input")
				continue
			}

			var dependentTasks []string
			if dependentTask != "" {
				dependentTasks = strings.Split(dependentTask, ",")
				for i := range dependentTasks {
					dependentTasks[i] = strings.TrimSpace(dependentTasks[i])
				}
			}

			tm.AddTask(task, linkedList, duration, priority, dependentTasks)
		case "complete":
			fmt.Println("Enter task index to complete:")
			indexStr, _ := reader.ReadString('\n')
			index, err := strconv.Atoi(strings.TrimSpace(indexStr))
			if err != nil {
				fmt.Println("Invalid index.")
				continue
			}
			if err := CompletedTask(&tm.tasks, index, saved); err != nil {
				fmt.Println("Error:", err)
			}
		case "delete":
			fmt.Println("Enter task index to Delete:")
			indexStr, _ := reader.ReadString('\n')
			index, err := strconv.Atoi(strings.TrimSpace(indexStr))
			if err != nil {
				fmt.Println("Invalid index.")
				continue
			}
			tm.DeleteTask(saved, index)
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
		case "suggestpriority":
			PrioritySuggestions(&tm.tasks)
		case "searchtask":
			fmt.Println("Enter task to search:")
			task, _ := reader.ReadString('\n')
			FindTask(strings.TrimSpace(task), linkedList)
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Unknown command.")
		}
	}
}
