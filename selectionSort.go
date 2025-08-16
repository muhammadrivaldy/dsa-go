package main

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		selectionIdx := i
		for j := (i + 1); j < len(arr); j++ {
			if arr[selectionIdx] > arr[j] {
				selectionIdx = j
			}
		}
		arr[selectionIdx], arr[i] = arr[i], arr[selectionIdx]
	}
	return arr
}
