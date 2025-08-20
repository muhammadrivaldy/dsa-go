package main

import (
	"fmt"
	"strings"
)

type singlyNode struct {
	Value int
	Next  *singlyNode
}

func singlyLinkedList(arr []int) {

	fmt.Println("Running singly linked list")

	head := &singlyNode{}
	node := head

	for _, i := range arr {
		temp := &singlyNode{Value: i}
		node.Next = temp
		node = temp
	}

	output := []string{}
	node = head.Next
	for node != nil {
		output = append(output, fmt.Sprint(node.Value))
		node = node.Next
	}

	fmt.Println("Completed to run singly linked list")
	fmt.Printf("The output is: [%v]\n", strings.Join(output, ", "))
}

type doublyNode struct {
	Value int
	Prev  *doublyNode
	Next  *doublyNode
}

func doublyLinkedList(arr []int) {

	fmt.Println("Running doublye linked list")

	head := &doublyNode{}
	node := head

	for _, i := range arr {

		temp := &doublyNode{Value: i}
		if node.Value > 0 {
			temp.Prev = node
		}

		node.Next = temp
		node = node.Next
	}

	output := []string{}

	node = head.Next
	isGoingBack := false
	for node != nil {
		if !isGoingBack {
			output = append(output, fmt.Sprint(node.Value))
			if node.Next == nil {
				isGoingBack = true
				node = node.Prev
				continue
			}
			node = node.Next
		} else {
			output = append(output, fmt.Sprint(node.Value))
			node = node.Prev
		}
	}

	fmt.Println("Completed to run doubly linked list")
	fmt.Printf("The output is: [%v]\n", strings.Join(output, ", "))
}

type circualNode struct {
	Value int
	Next  *circualNode
}

func circualLinkedList(arr []int) {

	fmt.Println("Running circual linked list")

	head := &circualNode{}
	node := head

	for _, i := range arr {
		temp := &circualNode{Value: i}
		node.Next = temp
		node = node.Next
	}

	node.Next = head.Next
	node = node.Next

	firstVal := node.Value
	repeat := 0
	output := []string{}

	for node != nil {

		if firstVal == node.Value {
			repeat++
			if repeat > 3 {
				break
			}
		}

		output = append(output, fmt.Sprint(node.Value))
		node = node.Next
	}

	fmt.Println("Completed to run circual linked list")
	fmt.Printf("The output is: [%v]\n", strings.Join(output, ", "))
}
