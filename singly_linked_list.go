package main

import (
	"fmt"
	"strings"
)

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

func singlyPrint(node *singlyNode) {

	output := []string{}

	for node != nil {
		output = append(output, fmt.Sprint(node.value))
		node = node.next
	}

	fmt.Printf("Finished, output: [%v]\n", strings.Join(output, ", "))
}
