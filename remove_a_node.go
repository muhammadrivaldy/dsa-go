package main

func removeNode(targetValue int, node *singlyNode) *singlyNode {
	if node.value == targetValue {
		return node.next
	}

	head := node
	for node.next != nil {
		if node.next.value == targetValue {
			node.next = node.next.next
			break
		}
		node = node.next
	}

	return head
}
