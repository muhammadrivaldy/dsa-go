package main

func linearSearch(arr []int, search int) int {
	index := -1
	for i, v := range arr {
		if search == v {
			index = i
			break
		}
	}
	return index
}
