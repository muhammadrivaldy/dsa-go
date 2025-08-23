package main

func binarySearch(arr []int, search int) int {
	idx := -1
	min, max := 0, len(arr)-1

	for (max - min) > 1 {
		mid := (max + min) / 2
		if arr[mid] == search {
			idx = mid
			break
		}

		if arr[mid] > search {
			max = mid
		} else {
			min = mid
		}
	}

	return idx
}
