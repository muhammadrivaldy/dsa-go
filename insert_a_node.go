package main

func insertNode(insertValue int, afterValue int, node *singlyNode) *singlyNode {
	head := node

	for node != nil {
		if node.value == afterValue {
			nextNode := node.next
			temp := &singlyNode{value: insertValue, next: nextNode}
			node.next = temp
			break
		}
		node = node.next
	}

	return head
}
