package main

func selectionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		smallIdx := i
		for j := i; j < len(arr); j++ {
			if arr[smallIdx] > arr[j] {
				smallIdx = j
			}
		}
		if smallIdx != i {
			arr[smallIdx], arr[i] = arr[i], arr[smallIdx]
		}
	}

	return arr
}
