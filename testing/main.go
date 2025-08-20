package main

import "fmt"

type tes struct {
	val int
}

func main() {
	a := &tes{val: 10}
	b := a
	c := b
	d := a
	b.val = 1
	c.val = 2
	temp := &tes{val: 3}
	d = temp
	fmt.Printf("%p %p %p\n", b, c, d)
	fmt.Println(a, b, c, d)

	g := 1
	h := &g
	i := h
	j := &g
	numTemp := 1
	j = &numTemp

	fmt.Println(h, i, j)

	x := &tes{val: 2}
	y := x
	temp = &tes{val: 15}
	y = temp
	fmt.Println(x, y, temp)
}
