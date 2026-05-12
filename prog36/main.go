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
func main() {
	root := buildTree()

	// Recursive
	var recursiveResult []int
	inorderRecursive(root, &recursiveResult)
	fmt.Println("Recursive:", recursiveResult)
}
func inorderRecursive(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	inorderRecursive(root.Left, result)  // 1. Visit left subtree
	*result = append(*result, root.Val)  // 2. Process current node
	inorderRecursive(root.Right, result) // 3. Visit right subtree
}
