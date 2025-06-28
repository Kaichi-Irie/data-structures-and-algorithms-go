package tree

import (
	"testing"
)

func TestBinaryTree(t *testing.T) {
	// Create a new binary tree
	tree := &BinaryTree{
		Root: 1,
		Left: &BinaryTree{
			Root: 2,
			Left: &BinaryTree{
				Root: 4,
			},
			Right: &BinaryTree{
				Root: 5,
			},
		},
		Right: &BinaryTree{
			Root: 3,
			Left: &BinaryTree{
				Root: 6,
			},
			Right: &BinaryTree{
				Root: 7,
			},
		},
	}

	if tree.Root != 1 {
		t.Errorf("Expected root to be 1, got %d", tree.Root)
	}
	if tree.Left.Root != 2 {
		t.Errorf("Expected left child to be 2, got %d", tree.Left.Root)
	}
	if tree.Right.Root != 3 {
		t.Errorf("Expected right child to be 3, got %d", tree.Right.Root)
	}
}
func TestDepth(t *testing.T) {
	// Create a new binary tree
	tree := &BinaryTree{
		Root: 1,
		Left: &BinaryTree{
			Root: 2,
			Left: &BinaryTree{
				Root: 4,
			},
			Right: &BinaryTree{
				Root: 5,
			},
		},
		Right: &BinaryTree{
			Root: 3,
			Left: &BinaryTree{
				Root: 6,
			},
			Right: &BinaryTree{
				Root: 7,
			},
		},
	}

	depth := tree.Depth()
	if depth != 3 {
		t.Errorf("Expected depth to be 3, got %d", depth)
	}
}

func TestNewBinarySearchTree(t *testing.T) {
	sortedArr := []int{1, 2, 3, 4, 5, 6, 7}
	bst := NewBinarySearchTree(sortedArr)

	if bst.Root != 4 {
		t.Errorf("Expected root to be 4, got %d", bst.Root)
	}
	if bst.Left.Root != 2 {
		t.Errorf("Expected left child to be 2, got %d", bst.Left.Root)
	}
	if bst.Right.Root != 6 {
		t.Errorf("Expected right child to be 6, got %d", bst.Right.Root)
	}
}
