package main

func selectionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		selectionIdx := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[selectionIdx] {
				selectionIdx = j
			}
		}
		arr[i], arr[selectionIdx] = arr[selectionIdx], arr[i]
	}

	return arr
}
