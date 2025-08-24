package main

func selectionSortLinkedList(node *singlyNode) *singlyNode {
	head := node
	first := head
	for first != nil {
		small := first
		check := first
		for check.next != nil {
			if small.value > check.next.value {
				small = check.next
			}
			check = check.next
		}
		first.value, small.value = small.value, first.value
		first = first.next
	}
	return head
}
