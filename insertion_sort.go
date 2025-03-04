package main

func insertionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {

		insertIndex := i
		currentValue := arr[i]

		for j := (i - 1); j >= 0; j-- {

			if arr[j] > currentValue {
				arr[(j + 1)] = arr[j]
				insertIndex = j
			} else {
				break
			}
		}

		arr[insertIndex] = currentValue
	}

	return arr
}
