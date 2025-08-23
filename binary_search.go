package main

// min 3
// max 4
// mid 3

func binarySearch(arr []int, search int) int {
	index := -1
	min, max := 0, len(arr)-1
	for {
		mid := (min + max) / 2
		if arr[mid] == search {
			index = mid
			break
		}

		if arr[mid] < search {
			min = mid
		} else if arr[mid] > search {
			max = mid
		}

		if max-min == 1 {
			break
		}
	}
	return index
}
