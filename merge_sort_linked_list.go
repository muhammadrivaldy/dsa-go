package main

func mergeSortLinkedList(node *singlyNode) *singlyNode {
	if node == nil || node.next == nil {
		return node
	}

	a, b := mergeSortLinkedListSplitIntoTwo(node)

	a = mergeSortLinkedList(a)
	b = mergeSortLinkedList(b)

	return mergeSortLinkedListSorted(a, b)
}

func mergeSortLinkedListSplitIntoTwo(node *singlyNode) (*singlyNode, *singlyNode) {
	if node == nil || node.next == nil {
		return node, nil
	}

	slow := node
	fast := node.next

	for fast != nil {
		fast = fast.next
		if fast != nil {
			slow = slow.next
			fast = fast.next
		}
	}

	mid := slow.next
	slow.next = nil

	return node, mid
}

func mergeSortLinkedListSorted(a *singlyNode, b *singlyNode) *singlyNode {
	if a == nil {
		return b
	} else if b == nil {
		return a
	}

	var output *singlyNode

	if a.value < b.value {
		output = a
		output.next = mergeSortLinkedListSorted(a.next, b)
	} else {
		output = b
		output.next = mergeSortLinkedListSorted(a, b.next)
	}

	return output
}
