package main

type singlyNode struct {
	value int
	next  *singlyNode
}

func singlyLinkedList(arr []int) *singlyNode {
	head := &singlyNode{}
	node := head

	for _, i := range arr {
		temp := &singlyNode{value: i}
		node.next = temp
		node = node.next
	}

	return head.next
}
