package main

func selectionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {

		selectIdx := i

		for j := (i + 1); j < len(arr); j++ {

			if arr[j] < arr[selectIdx] {
				selectIdx = j
			}
		}

		arr[i], arr[selectIdx] = arr[selectIdx], arr[i]
	}

	return arr
}
