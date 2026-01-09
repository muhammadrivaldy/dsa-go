package main

func selectionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		selcIdx := i
		for j := i; j < len(arr); j++ {
			if arr[selcIdx] > arr[j] {
				selcIdx = j
			}
		}
		arr[selcIdx], arr[i] = arr[i], arr[selcIdx]
	}

	return arr
}
