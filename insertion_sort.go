package main

func insertionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {

		insertIdx := i
		insertVal := arr[i]

		for j := (i - 1); j >= 0; j-- {
			if arr[j] > insertVal {
				arr[j+1] = arr[j]
				insertIdx = j
			}
		}

		arr[insertIdx] = insertVal
	}

	return arr
}
