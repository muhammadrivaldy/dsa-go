package main

func selectionSortLinkedList(node *singlyNode) *singlyNode {

	head := node

	for node != nil {

		child := node.next
		target := node

		for child != nil {

			if target.value > child.value {
				target = child
			}
			child = child.next
		}

		node.value, target.value = target.value, node.value
		node = node.next
	}

	return head
}
