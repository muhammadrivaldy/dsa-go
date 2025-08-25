package main

import (
	"fmt"
)

// Node structure
type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	Height int
}

// Utility: get height of node
func height(n *Node) int {
	if n == nil {
		return 0
	}
	return n.Height
}

// Utility: max of two ints
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Get balance factor
func getBalance(n *Node) int {
	if n == nil {
		return 0
	}
	return height(n.Left) - height(n.Right)
}

// Right rotation
func rightRotate(y *Node) *Node {
	x := y.Left
	T2 := x.Right

	// Perform rotation
	x.Right = y
	y.Left = T2

	// Update heights
	y.Height = max(height(y.Left), height(y.Right)) + 1
	x.Height = max(height(x.Left), height(x.Right)) + 1

	return x
}

// Left rotation
func leftRotate(x *Node) *Node {
	y := x.Right
	T2 := y.Left

	// Perform rotation
	y.Left = x
	x.Right = T2

	// Update heights
	x.Height = max(height(x.Left), height(x.Right)) + 1
	y.Height = max(height(y.Left), height(y.Right)) + 1

	return y
}

// Insert into AVL tree
func insert(node *Node, key int) *Node {
	// 1. Normal BST insert
	if node == nil {
		return &Node{Value: key, Height: 1}
	}
	if key < node.Value {
		node.Left = insert(node.Left, key)
	} else if key > node.Value {
		node.Right = insert(node.Right, key)
	} else {
		// Equal keys not allowed
		return node
	}

	// 2. Update height
	node.Height = 1 + max(height(node.Left), height(node.Right))

	// 3. Get balance factor
	balance := getBalance(node)

	// 4. If unbalanced → rotations
	// Left Left Case
	if balance > 1 && key < node.Left.Value {
		return rightRotate(node)
	}

	// Right Right Case
	if balance < -1 && key > node.Right.Value {
		return leftRotate(node)
	}

	// Left Right Case
	if balance > 1 && key > node.Left.Value {
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}

	// Right Left Case
	if balance < -1 && key < node.Right.Value {
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}

	return node
}

// Inorder traversal
func inorder(root *Node) {
	if root != nil {
		inorder(root.Left)
		fmt.Print(root.Value, " ")
		inorder(root.Right)
	}
}

func main() {
	var root *Node
	values := []int{5, 2, 6, 3, 4}

	for _, v := range values {
		root = insert(root, v)
	}

	fmt.Print("Inorder traversal of AVL tree: ")
	inorder(root)
	fmt.Println()
}
