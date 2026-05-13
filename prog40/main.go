package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree() *TreeNode {
	return &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   6,
			Left:  &TreeNode{Val: 5},
			Right: &TreeNode{Val: 7},
		},
	}
}

func postorderRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	result = append(result, postorderRecursive(root.Left)...)
	result = append(result, postorderRecursive(root.Right)...)
	result = append(result, root.Val)
	return result
}

func main() {
	root := buildTree()
	result := postorderRecursive(root)
	fmt.Println(result)
}
