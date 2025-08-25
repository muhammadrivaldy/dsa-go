package main

func linearSearch(arr []int, search int) int {

	for i, v := range arr {
		if v == search {
			return i
		}
	}

	return -1
}
