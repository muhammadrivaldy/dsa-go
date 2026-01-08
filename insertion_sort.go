package main

func insertionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		insertionIdx := i
		insertionVal := arr[i]
		for j := i; j >= 0; j-- {
			if arr[j] > insertionVal {
				arr[insertionIdx] = arr[j]
				insertionIdx = j
			}
		}
		arr[insertionIdx] = insertionVal
	}

	return arr
}
