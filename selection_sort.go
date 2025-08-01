package main

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minIdx := i
		for j := (i + 1); j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[minIdx], arr[i] = arr[i], arr[minIdx]
	}
	return arr
}
