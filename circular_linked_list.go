package main

type circularNode struct {
	value int
	next  *circularNode
}

func circularLinkedList(arr []int) *circularNode {

	head := &circularNode{}
	node := head

	for _, i := range arr {
		temp := &circularNode{value: i}
		node.next = temp
		node = node.next
	}

	node.next = head.next

	return head.next
}
