package main

func insertNode(value, after int, node *singlyNode) *singlyNode {

	head := node

	for node != nil {
		if node.value == after {
			temp := node.next
			node.next = &singlyNode{value: value, next: temp}
			return head
		}
		node = node.next
	}

	return head
}
