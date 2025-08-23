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
	fmt.Println("Starting run singly linked list")
	head := &singlyNode{}
	node := head

	for _, i := range arr {
		temp := &singlyNode{value: i}
		node.next = temp
		node = node.next
	}

	valStrings := []string{}
	node = head.next
	for node != nil {
		valStrings = append(valStrings, fmt.Sprint(node.value))
		node = node.next
	}

	fmt.Printf("Output: [%v]\n", strings.Join(valStrings, ", "))
	fmt.Println("Finished run singly")
}

type doublyNode struct {
	value int
	prev  *doublyNode
	next  *doublyNode
}

func doublyLinkedList(arr []int) {
	fmt.Println("Starting run doubly linked list")
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

	valStrings := []string{}
	node = head.next
	isPrev := false
	for node != nil {
		valStrings = append(valStrings, fmt.Sprint(node.value))
		if !isPrev {
			if node.next == nil {
				node = node.prev
				isPrev = true
				continue
			}
			node = node.next
		} else {
			node = node.prev
		}
	}

	fmt.Printf("Output: [%v]\n", strings.Join(valStrings, ", "))
	fmt.Println("Finished run doubly")
}

type circularNode struct {
	value int
	next  *circularNode
}

func circularLinkedList(arr []int) {
	fmt.Println("Starting run circular linked list")
	head := &circularNode{}
	node := head

	for _, i := range arr {
		temp := &circularNode{value: i}
		node.next = temp
		node = temp
	}

	valStrings := []string{}
	node.next = head.next
	node = head.next
	repeat := 0
	target := 10
	for node != nil {
		valStrings = append(valStrings, fmt.Sprint(node.value))
		node = node.next
		if node.value == arr[0] {
			repeat++
			if repeat == target {
				break
			}
		}
	}

	fmt.Printf("Output: [%v]\n", strings.Join(valStrings, ", "))
	fmt.Println("Finished run circular")
}
