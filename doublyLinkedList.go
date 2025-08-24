package main

type doublyNode struct {
	value int
	prev  *doublyNode
	next  *doublyNode
}

func doublyLinkedList(arr []int) *doublyNode {
	head := &doublyNode{}
	node := head

	for _, i := range arr {
		temp := &doublyNode{value: i}
		if node.value > 0 {
			temp.prev = node
		}

		node.next = temp
		node = node.next
	}

	return head.next
}
