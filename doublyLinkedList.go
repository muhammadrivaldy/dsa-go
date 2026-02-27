package main

type doublyNode struct {
	value int
	next  *doublyNode
	prev  *doublyNode
}

func doublyLinkedList(arr []int) *doublyNode {

	head := &doublyNode{}
	node := head

	for _, i := range arr {
		temp := &doublyNode{value: i}
		node.next = temp
		temp.prev = node
		node = temp
	}

	return head.next
}
