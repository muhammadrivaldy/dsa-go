package main

func insertionSort(arr []int) []int {

	// input [3, 2, 4, 1]

	for i := 0; i < len(arr); i++ {

		insertIdx := i
		currentValue := arr[i]

		for j := (i - 1); j >= 0; j-- {
			if arr[j] > currentValue {
				arr[j+1] = arr[j]
				insertIdx = j
			} else {
				break
			}
		}

		arr[insertIdx] = currentValue
	}

	return arr
}
