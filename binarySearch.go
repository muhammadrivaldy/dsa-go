package main

func binarySearch(arr []int, val int) int {
	min, max := 0, len(arr)-1
	for {
		mid := (min + max) / 2
		if arr[mid] == val {
			return mid
		} else if arr[mid] < val {
			min = mid
		} else if arr[mid] > val {
			max = mid
		} else {
			return -1
		}
	}
}
