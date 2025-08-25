package main

func bubbleSortLinkedList(node *singlyNode) *singlyNode {

	head := node

	for {

		node = head
		isSwap := false

		for node.next != nil {

			if node.next.value < node.value {
				node.next.value, node.value = node.value, node.next.value
				isSwap = true
			}
			node = node.next
		}

		if !isSwap {
			break
		}
	}

	return head
}
