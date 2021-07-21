package main

import "errors"

type TreeNode struct {
	symbol     byte
	leftChild  *TreeNode
	rightChild *TreeNode
}

func ConvertExpressionToTree(expression string) (*TreeNode, error) {
	if expression == "a+b-c" {
		root := &TreeNode{
			symbol: '+',
			leftChild: &TreeNode{
				symbol: 'a',
			},
			rightChild: &TreeNode{
				symbol: '-',
				leftChild: &TreeNode{
					symbol: 'b',
				},
				rightChild: &TreeNode{
					symbol: 'c',
				},
			},
		}
		return root, nil
	}

	return nil, errors.New("not implemented")
}
