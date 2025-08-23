package main

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		target := i
		value := arr[i]
		for j := (i - 1); j >= 0; j-- {
			if arr[j] > value {
				arr[j+1] = arr[j]
				target = j
			}
		}
		arr[target] = value
	}
	return arr
}
