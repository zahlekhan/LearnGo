package main

import "errors"

type TreeNode struct {
	symbol     byte
	leftChild  *TreeNode
	rightChild *TreeNode
}

func (t *TreeNode) PostOrder() string {
	leftChildPostOrder := ""
	if t.leftChild != nil {
		leftChildPostOrder = t.leftChild.PostOrder()
	}

	rightChildPostOrder := ""
	if t.rightChild != nil {
		rightChildPostOrder = t.rightChild.PostOrder()
	}

	return leftChildPostOrder + rightChildPostOrder + string(t.symbol)
}

func (t *TreeNode) PreOrder() string {
	leftChildPreOrder := ""
	if t.leftChild != nil {
		leftChildPreOrder = t.leftChild.PreOrder()
	}

	rightChildPreOrder := ""
	if t.rightChild != nil {
		rightChildPreOrder = t.rightChild.PreOrder()
	}

	return string(t.symbol) + leftChildPreOrder + rightChildPreOrder
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
