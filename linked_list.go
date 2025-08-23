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

	fmt.Println("Running singly linked list")

	head := &singlyNode{}
	node := head
	for _, i := range arr {
		temp := &singlyNode{value: i}
		node.next = temp
		node = node.next
	}

	output := []string{}

	node = head.next
	for node != nil {
		output = append(output, fmt.Sprint(node.value))
		node = node.next
	}

	fmt.Printf("Finished, output: [%v]\n", strings.Join(output, ", "))
}

type doublyNode struct {
	value int
	prev  *doublyNode
	next  *doublyNode
}

func doublyLinkedList(arr []int) {

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

	output := []string{}
	isGoingBack := false

	node = head.next
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

type circularNode struct {
	value int
	next  *circularNode
}

func circularLinkedList(arr []int) {

	fmt.Println("Running circular linked list")

	head := &circularNode{}
	node := head
	for _, i := range arr {
		temp := &circularNode{value: i}
		node.next = temp
		node = node.next
	}
	node.next = head.next
	node = head.next

	output := []string{}
	repeat := 0
	target := 2

	for node != nil {
		if node.value == arr[0] {
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
