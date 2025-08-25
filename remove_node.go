package main

func removeNode(value int, node *singlyNode) *singlyNode {

	head := node

	if node.value == value {
		return node.next
	}

	for node.next != nil {
		if node.next.value == value {
			node.next = node.next.next
			break
		}
		node = node.next
	}

	return head
}
