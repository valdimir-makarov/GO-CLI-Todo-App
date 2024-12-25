package main

type item struct {
	Task string
}
type List []item
type treeNode struct {
	state      List
	parentNode *treeNode
	childNode  []*treeNode
}

var root = &treeNode{
	state:      List{},
	parentNode: nil,
	childNode:  []*treeNode{},
}

func addTask(l *List, t *item) {

}

func main() {

}
