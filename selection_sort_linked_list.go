package main

func selectionSortLinkedList(node *singlyNode) *singlyNode {
	head := node
	parent := head
	for parent != nil {
		child := parent.next
		selection := parent
		for child != nil {
			if selection.value > child.value {
				selection = child
			}
			child = child.next
		}
		parent.value, selection.value = selection.value, parent.value
		parent = parent.next
	}
	return head
}
