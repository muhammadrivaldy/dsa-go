package main

func selectionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		selectIdx := i
		for j := i; j < len(arr); j++ {
			if arr[selectIdx] > arr[j] {
				selectIdx = j
			}
		}
		arr[i], arr[selectIdx] = arr[selectIdx], arr[i]
	}

	return arr
}
