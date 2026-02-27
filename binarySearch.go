package main

func binarySearch(arr []int, search int) int {

	min, max := 0, len(arr)-1

	for min <= max {
		mid := min + (max-min)/2
		if arr[mid] == search {
			return mid
		}

		if arr[mid] < search {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}

	return 0
}
