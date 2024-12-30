// package main

// import (
// 	"fmt"
// 	"time"
// )

// type item struct {
// 	Task        string
// 	Done        bool
// 	CreatedAt   time.Time
// 	CompletedAt time.Time
// }

// type Tree struct {
// 	Task        string
// 	ParentNode  *Tree
// 	ChildNodes  []*Tree
// 	CurrentNode string
// }

// type Stack struct {
// 	trashedItems []string
// }

// var stack Stack

// type Saved struct {
// 	savedItems []string
// }
// type List []item

// func (tree *Tree) AddTask(l *List, task string) {
// 	t := item{
// 		Task:        task,
// 		Done:        false,
// 		CreatedAt:   time.Now(),
// 		CompletedAt: time.Time{},
// 	}
// 	*l = append(*l, t)
// 	NewTreeVlaue := &Tree{
// 		Task:       task,
// 		ParentNode: tree,
// 		ChildNodes: []*Tree{},
// 	}
// 	tree.ChildNodes = append(tree.ChildNodes, NewTreeVlaue)

// }
// func (tree *Tree) undo() *Tree {

// 	if tree.ParentNode != nil {
// 		fmt.Println("Undo: Moving to parent task:", tree.ParentNode.Task)
// 		return tree.ParentNode
// 	}
// 	fmt.Println("Undo: Already at the root task.")
// 	return tree

// }
// func PrintTree(node *Tree, depth int) {
// 	fmt.Printf("%s- %s\n", string(' '+depth*2), node.Task)
// 	for _, child := range node.ChildNodes {
// 		PrintTree(child, depth+1)
// 	}
// }

// func CompletedTask(l *List, i int, saved *Saved) error {

// 	ls := *l
// 	if i <= 0 || i > len(ls) {
// 		return fmt.Errorf("item dose not exist")
// 	}
// 	ls[i-1].Done = true
// 	if ls[i-1].Done {

// 		SaveTask(&ls[i-1], saved)
// 	}
// 	ls[i-1].CompletedAt = time.Now()
// 	return nil

// }
// func SaveTask(task *item, saved *Saved) {

// 	saved.savedItems = append(saved.savedItems, task.Task)
// 	fmt.Println("Task saved:", task.Task)
// }
// func DeleteTask(saved *Saved) {
// 	TrashItems(saved, &stack)
// 	saved.savedItems = nil
// 	fmt.Println("the Saved Task was Deleted", saved.savedItems)
// }
// func TrashItems(saved *Saved, s *Stack) {
// 	for _, task := range saved.savedItems {
// 		s.trashedItems = append(s.trashedItems, task)
// 	}
// 	fmt.Println("Tasks moved to trash:", s.trashedItems)

// }

// // func UndoRedo() { //i guess this the most complex function till now
// // 	// so the function actually works   undo in o(1) and redo in o(n)
// // 	// in the future i am gonna use AVL trees that will make the redo functionality o(nlongn);
// // 	// package main

// // 	// import "fmt"

// // 	// type Tree struct {
// // 	// 	ParentNode  *Tree
// // 	// 	ChildNode   []*Tree
// // 	// 	CurrentNode string
// // 	// }

// // 	// func (tree *Tree) AddChild(task string) *Tree {
// // 	// 	newNode := &Tree{
// // 	// 		ParentNode:  tree,
// // 	// 		ChildNode:   []*Tree{},
// // 	// 		CurrentNode: task,
// // 	// 	}
// // 	// 	tree.ChildNode = append(tree.ChildNode, newNode)
// // 	// 	return newNode
// // 	// }

// // 	// func (tree *Tree) Undo() *Tree {
// // 	// 	// Option 2: Move to Parent
// // 	// 	if tree.ParentNode != nil {
// // 	// 		return tree.ParentNode
// // 	// 	}

// // 	// 	// If already at root, stay in the current node
// // 	// 	return tree
// // 	// }

// // 	// func PrintTree(node *Tree, depth int) {
// // 	// 	fmt.Printf("%s%s\n", string(' '+depth*2), node.CurrentNode)
// // 	// 	for _, child := range node.ChildNode {
// // 	// 		PrintTree(child, depth+1)
// // 	// 	}
// // 	// }

// // 	// func main() {
// // 	// 	root := &Tree{
// // 	// 		ParentNode:  nil,
// // 	// 		ChildNode:   []*Tree{},
// // 	// 		CurrentNode: "task",
// // 	// 	}

// // 	// 	current := root
// // 	// 	current = current.AddChild("task1")
// // 	// 	current = current.AddChild("task2")
// // 	// 	current = current.AddChild("task3")
// // 	// 	current = current.AddChild("Task4")

// // 	// 	// Undo
// // 	// 	current = current.Undo()
// // 	// 		current = current.Undo()

// // 	// 	PrintTree(root, 0)
// // 	// }

// // }
