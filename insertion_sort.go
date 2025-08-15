package main

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		tIdx := i
		actualVal := arr[i]
		for j := (i - 1); j >= 0; j-- {
			if actualVal < arr[j] {
				arr[j+1] = arr[j]
				tIdx = j
			}
		}
		arr[tIdx] = actualVal
	}
	return arr
}
