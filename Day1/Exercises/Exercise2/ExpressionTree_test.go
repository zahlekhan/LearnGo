package main

import (
	"testing"
)

func TestConvertExpressionToTree(t *testing.T) {
	t.Run("a+b-c", func(t *testing.T) {
		expression := "a+b-c"
		got, _ := ConvertExpressionToTree(expression)

		want := &TreeNode{
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

		if !equalTree(t, got, want) {
			t.Errorf("trees are different for %s", expression)
		}
	})

	t.Run("d+b-c", func(t *testing.T) {
		expression := "d+b-c"
		_, err := ConvertExpressionToTree(expression)

		if err == nil {
			t.Fatalf("expected an error")
		}

		want := "not implemented"
		if err.Error() != want {
			t.Errorf("got %q want %q", err.Error(), want)
		}
	})
}

func TestTreeNode_PostOrder(t *testing.T) {
	t.Run("a+b-c", func(t *testing.T) {
		expression := "a+b-c"
		root, _ := ConvertExpressionToTree(expression)

		got := root.PostOrder()
		want := "abc-+"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestTreeNode_PreOrder(t *testing.T) {
	t.Run("a+b-c", func(t *testing.T) {
		expression := "a+b-c"
		root, _ := ConvertExpressionToTree(expression)

		got := root.PreOrder()
		want := "+a-bc"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func equalTree(t testing.TB, t1 *TreeNode, t2 *TreeNode) bool {
	t.Helper()
	if t1 == nil && t2 == nil {
		return true
	}

	if (t1 == nil && t2 != nil) || (t1 != nil && t2 == nil) || (t1.symbol != t2.symbol) {
		return false
	}

	return equalTree(t, t1.leftChild, t2.leftChild) && equalTree(t, t1.rightChild, t2.rightChild)
}
