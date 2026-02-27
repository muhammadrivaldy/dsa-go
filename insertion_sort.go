package main

func insertionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		idx := i
		val := arr[i]
		for j := i; j >= 0; j-- {
			if arr[j] > val {
				arr[j+1] = arr[j]
				idx = j
			}
		}
		arr[idx] = val
	}

	return arr
}
