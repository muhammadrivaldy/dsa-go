package main

func linearSearch(arr []int, val int) (idx int) {
	for i, j := range arr {
		if val == j {
			return i
		}
	}
	return -1
}
