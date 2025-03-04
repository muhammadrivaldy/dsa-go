package main

func selectionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {

		minIndex := i

		for j := i + 1; j < len(arr); j++ {

			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}

		if minIndex != i {
			arr[minIndex], arr[i] = arr[i], arr[minIndex]
		}
	}

	return arr
}
