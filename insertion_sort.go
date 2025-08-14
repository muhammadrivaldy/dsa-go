package main

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		insertIdx := i
		currentVal := arr[i]
		for j := (i - 1); j >= 0; j-- {
			if arr[j] > currentVal {
				arr[j+1] = arr[j]
				insertIdx = j
			}
		}
		arr[insertIdx] = currentVal
	}
	return arr
}
