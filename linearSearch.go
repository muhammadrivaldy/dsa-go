package main

func linearSearch(arr []int, search int) int {
	for idx, i := range arr {
		if i == search {
			return idx
		}
	}
	return 0
}
