package main

import (
	"fmt"
	"strings"
)

type singlyNode struct {
	value int
	next  *singlyNode
}

func singlyLinkedList(arr []int) {
	output := []string{}
	fmt.Println("Running singly linked list")

	head := &singlyNode{}
	node := head

	for _, i := range arr {
		temp := &singlyNode{value: i}
		node.next = temp
		node = node.next
	}

	node = head.next
	for node != nil {
		output = append(output, fmt.Sprint(node.value))
		node = node.next
	}

	fmt.Printf("Finished running singly linked list, output: [%v]\n", strings.Join(output, ", "))
}

type doublyNode struct {
	value int
	prev  *doublyNode
	next  *doublyNode
}

func doublyLinkedList(arr []int) {
	output := []string{}
	fmt.Println("Running doubly linked list")

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

	node = head.next
	isPrev := false
	for node != nil {
		output = append(output, fmt.Sprint(node.value))
		if !isPrev {
			if node.next == nil {
				isPrev = true
				node = node.prev
				continue
			}
			node = node.next
		} else {
			node = node.prev
		}
	}

	fmt.Printf("Finished running doubly linked list, output: [%v]\n", strings.Join(output, ", "))
}

type circularNode struct {
	value int
	next  *circularNode
}

func circularLinkedList(arr []int) {
	output := []string{}
	fmt.Println("Running circular linked list")

	head := &circularNode{}
	node := head

	for _, i := range arr {
		temp := &circularNode{value: i}
		node.next = temp
		node = node.next
	}

	node.next = head.next
	repeat, target := 0, 2
	node = head.next
	for node != nil {
		if node.value == arr[0] {
			repeat++
			if repeat > target {
				break
			}
		}
		output = append(output, fmt.Sprint(node.value))
		node = node.next
	}

	fmt.Printf("Finished running circular doubly linked list, output: [%v]\n", strings.Join(output, ", "))
}
