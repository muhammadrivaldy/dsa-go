package main

func insertionSortLinkedList(node *singlyNode) *singlyNode {
	if node == nil || node.next == nil {
		return node
	}

	var sorted *singlyNode
	current := node

	for current != nil {
		temp := current.next
		sorted = insertionSortLinkedListSorted(sorted, current)
		current = temp
	}

	return sorted
}

func insertionSortLinkedListSorted(sorted *singlyNode, current *singlyNode) *singlyNode {
	if sorted == nil || sorted.value >= current.value {
		current.next = sorted
		return current
	}

	temp := sorted
	for temp.next != nil && temp.next.value <= current.value {
		temp = temp.next
	}

	current.next = temp.next
	temp.next = current

	return sorted
}
