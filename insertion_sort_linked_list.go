package main

func insertionSortLinkedList(node *singlyNode) *singlyNode {

	if node == nil || node.next == nil {
		return node
	}

	var sorted *singlyNode
	for node != nil {
		temp := node.next
		sorted = insertionSortLinkedListSorted(sorted, node)
		node = temp
	}

	return sorted
}

func insertionSortLinkedListSorted(sorted, unsorted *singlyNode) *singlyNode {

	if sorted == nil || unsorted.value < sorted.value {
		unsorted.next = sorted
		return unsorted
	}

	temp := sorted
	if temp.next != nil && temp.next.value < unsorted.value {
		temp = temp.next
	}

	unsorted.next = temp.next
	temp.next = unsorted

	return sorted
}
