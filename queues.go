package main

import "fmt"

type queue struct {
	values []int
}

func (q *queue) print() {
	fmt.Printf("Output: %v\n", q.values)
}

func (q *queue) enqueue(val int) {
	fmt.Println("Enqueue:", val)
	q.values = append(q.values, val)
}

func (q *queue) dequeue() int {
	fmt.Println("Dequeue")
	if len(q.values) == 0 {
		return -1
	}

	value := q.values[0]
	q.values = q.values[1:]
	return value
}
