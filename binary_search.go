package main

func binarySearch(arr []int, search int) int {

	min, max := 0, len(arr)-1

	for (max - min) > 1 {

		mid := (max + min) / 2
		if arr[mid] == search {
			return mid
		}

		if arr[mid] < search {
			min = mid
		} else {
			max = mid
		}
	}

	return -1
}
