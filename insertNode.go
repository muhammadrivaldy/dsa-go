package main

func insertNode(value, after int, node *singlyNode) *singlyNode {
	head := node
	for node != nil {
		if node.value == after {
			if node.next != nil {
				temp := node.next
				node.next = &singlyNode{value: value, next: temp}
			} else {
				node.next = &singlyNode{value: value}
			}
			break
		}
		node = node.next
	}
	return head
}
