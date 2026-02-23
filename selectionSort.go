package main

func selectionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		selectIdx := i
		for j := (i + 1); j < len(arr); j++ {
			if arr[selectIdx] > arr[j] {
				selectIdx = j
			}
		}

		arr[selectIdx], arr[i] = arr[i], arr[selectIdx]
	}

	return arr
}
