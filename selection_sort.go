package main

func selectionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		lessIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[lessIdx] {
				lessIdx = j
			}
		}
		if lessIdx != i {
			arr[i], arr[lessIdx] = arr[lessIdx], arr[i]
		}
	}

	return arr
}
