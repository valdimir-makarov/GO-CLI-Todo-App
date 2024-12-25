package main

import (
	"fmt"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}
type Stack struct {
	trashedItems []string
}

var stack Stack

type Saved struct {
	savedItems []string
}
type List []item

func AddTask(l *List, task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*l = append(*l, t)

}
func CompletedTask(l *List, i int, saved *Saved) error {

	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("item dose not exist")
	}
	ls[i-1].Done = true
	if ls[i-1].Done {

		SaveTask(&ls[i-1], saved)
	}
	ls[i-1].CompletedAt = time.Now()
	return nil

}
func SaveTask(task *item, saved *Saved) {

	saved.savedItems = append(saved.savedItems, task.Task)
	fmt.Println("Task saved:", task.Task)
}
func DeleteTask(saved *Saved) {
	TrashItems(saved, &stack)
	saved.savedItems = nil
	fmt.Println("the Saved Task was Deleted", saved.savedItems)
}
func TrashItems(saved *Saved, s *Stack) {
	for _, task := range saved.savedItems {
		s.trashedItems = append(s.trashedItems, task)
	}
	fmt.Println("Tasks moved to trash:", s.trashedItems)

}
func UndoRedo (){//i guess this the most complex function till now 
	// so the function actually works   undo in o(1) and redo in o(n)
// in the future i am gonna use AVL trees that will make the redo functionality o(nlongn);





}