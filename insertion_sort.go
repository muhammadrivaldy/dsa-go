package main

func insertionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {

		insertIdx := i
		copyValue := arr[i]

		for j := (i - 1); j >= 0; j-- {

			if arr[j] > copyValue {
				arr[(j + 1)] = arr[j]
				insertIdx = j
			} else {
				break
			}
		}

		arr[insertIdx] = copyValue
	}

	return arr
}
