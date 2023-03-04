package main

import "fmt"

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func main() {
	var tr = &TreeNode{
		HasToy: true,
		Left: &TreeNode{
			HasToy: true,
			Left: &TreeNode{
				HasToy: true,
				Left:   nil,
				Right:  nil,
			},
			Right: &TreeNode{
				HasToy: false,
				Left:   nil,
				Right:  nil,
			},
		},
		Right: &TreeNode{
			HasToy: true,
			Left: &TreeNode{
				HasToy: false,
				Left:   nil,
				Right:  nil,
			},
			Right: &TreeNode{
				HasToy: true,
				Left:   nil,
				Right:  nil,
			},
		},
	}
	res := areToysBalanced(tr)
	fmt.Println(res)
}

func areToysBalanced(root *TreeNode) bool {
	var left, right int = 0, 0
	left = GetTreeNodeNum(root.Left)
	right = GetTreeNodeNum(root.Right)
	if left == right {
		return true
	} else {
		return false
	}
}

func GetTreeNodeNum(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.HasToy == true {
		return GetTreeNodeNum(root.Left) + GetTreeNodeNum(root.Right) + 1
	}
	return 0
}
