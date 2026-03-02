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
		temp.prev = node
		node.next = temp
		node = temp
	}

	return head.next
}
