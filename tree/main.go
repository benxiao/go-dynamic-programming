package main

import "fmt"

type TreeNode struct {
	Key   string
	Value string
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Value string
	Prev *TreeNode
	Next *TreeNode
}

//func add(lst *ListNode)
func (tree *TreeNode) Display() string {
	if tree == nil {
		return "nil"
	}
	return fmt.Sprintf("(%s, %s, %s)", tree.Key,
		tree.Left.Display(), tree.Right.Display())
}

func (tree *TreeNode)Insert(key string, value string) *TreeNode{
	if tree == nil {
		return &TreeNode{Key:key, Value:value, Left:nil, Right:nil}
	}
	if key < tree.Key {
		tree.Left = tree.Left.Insert(key, value)
	} else {
		tree.Right = tree.Right.Insert(key, value)
	}
	return tree
}



func main() {
	var tree *TreeNode;
	tree = tree.Insert("d", "a")
	fmt.Println(tree.Display())
}