package main

func linearSearch(arr []int, search int) int {
	idx := -1
	for i, v := range arr {
		if v == search {
			idx = i
			break
		}
	}
	return idx
}
