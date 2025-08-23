package main

import (
	"fmt"
	"strings"
)

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

func circularPrint(node *circularNode) {

	output := []string{}
	repeat := 0
	target := 2
	headValue := node.value

	for node != nil {
		if node.value == headValue {
			if repeat >= target {
				break
			}
			repeat++
		}

		output = append(output, fmt.Sprint(node.value))
		node = node.next
	}

	fmt.Printf("Finished, output: [%v]\n", strings.Join(output, ", "))
}
