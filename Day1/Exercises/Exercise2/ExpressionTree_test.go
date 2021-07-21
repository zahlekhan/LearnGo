package main

import "testing"

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
