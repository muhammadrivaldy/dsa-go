package main

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		idx := i
		for j := (i + 1); j < len(arr); j++ {
			if arr[j] < arr[idx] {
				idx = j
			}
		}
		arr[idx], arr[i] = arr[i], arr[idx]
	}
	return arr
}
