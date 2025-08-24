package main

func removeNode(target int, node *singlyNode) *singlyNode {
	head := node
	var prev *singlyNode
	for node != nil {
		if node.value == target {
			prev.next = node.next
			break
		}
		prev = node
		node = node.next
	}
	return head
}
