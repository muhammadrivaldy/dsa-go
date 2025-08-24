package main

import "fmt"

func getMax(arr []int) int {
	max := 0
	for _, val := range arr {
		if val > max {
			max = val
		}
	}

	return max
}

func singlyPrint(node *singlyNode) {
	output := []int{}
	for node != nil {
		output = append(output, node.value)
		node = node.next
	}
	fmt.Printf("Output: %v\n", output)
}

func doublyPrint(node *doublyNode) {
	output := []int{}
	isGoingBack := false
	for node != nil {
		output = append(output, node.value)
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
	fmt.Printf("Output: %v\n", output)
}

func circularPrint(node *circularNode) {
	output := []int{}
	repeat, target := 0, 2
	firstVal := node.value
	for node != nil {
		if node.value == firstVal {
			repeat++
			if repeat > target {
				break
			}
		}

		output = append(output, node.value)
		node = node.next
	}
	fmt.Printf("Output: %v\n", output)
}
