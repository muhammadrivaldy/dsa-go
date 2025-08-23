package main

import (
	"fmt"
	"strings"
)

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

func doublyPrint(node *doublyNode) {

	output := []string{}
	isGoingBack := false

	for node != nil {
		output = append(output, fmt.Sprint(node.value))
		if !isGoingBack {
			if node.next == nil {
				isGoingBack = true
				node = node.prev
				continue
			}
			node = node.next
		} else {
			node = node.prev
		}
	}

	fmt.Printf("Finished, output: [%v]\n", strings.Join(output, ", "))
}
