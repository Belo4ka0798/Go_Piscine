package main

import (
	"fmt"
)

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func main() {
	// создаем бинарное дерево
	root := &TreeNode{HasToy: true}
	root.Left = &TreeNode{HasToy: false}
	root.Right = &TreeNode{HasToy: true}
	root.Left.Left = &TreeNode{HasToy: false}
	root.Left.Right = &TreeNode{HasToy: true}
	root.Right.Left = &TreeNode{HasToy: false}
	root.Right.Right = &TreeNode{HasToy: true}

	// производим обход змейкой
	unrollGarland(root)
}

func unrollGarland(root *TreeNode) {
	if root == nil {
		return
	}
	var leftToRight bool = true
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		stackLen := len(stack)
		for i := 0; i < stackLen; i++ {
			top := stack[i]
			fmt.Print(top.HasToy, " ")
			if leftToRight {
				if top.Left != nil {
					stack = append(stack, top.Left)
				}
				if top.Right != nil {
					stack = append(stack, top.Right)
				}
			} else {
				if top.Right != nil {
					stack = append(stack, top.Right)
				}
				if top.Left != nil {
					stack = append(stack, top.Left)
				}
			}
		}
		leftToRight = !leftToRight
		stack = stack[stackLen:]
		fmt.Println()
	}
}
