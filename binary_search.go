package main

func binarySearch(arr []int, target int) int {

	min, max := 0, len(arr)-1

	for i := 0; i < len(arr); i++ {
		mid := min + (max-min)/2
		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}

	return 0
}
