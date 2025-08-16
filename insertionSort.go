package main

func insertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		insertionIdx := i
		insertionVal := arr[i]
		for j := (i - 1); j >= 0; j-- {
			if arr[j] > insertionVal {
				arr[j+1] = arr[j]
				insertionIdx = j
			}
		}
		arr[insertionIdx] = insertionVal
	}
	return arr
}
