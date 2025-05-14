package main

func insertionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {

		currentValue := arr[i]
		targetedIndex := i

		for j := (i - 1); j >= 0; j-- {

			if arr[j] > currentValue {
				arr[(j + 1)] = arr[j]
				targetedIndex = j
			} else {
				break
			}
		}

		arr[targetedIndex] = currentValue
	}

	return arr
}
