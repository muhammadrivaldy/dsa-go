package main

func insertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		targeteIdx := i
		currentVal := arr[i]
		for j := i - 1; j >= 0; j-- {
			if arr[j] > currentVal {
				arr[j+1] = arr[j]
				targeteIdx = j
			}
		}
		arr[targeteIdx] = currentVal
	}
	return arr
}
