package main

func insertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		targetedIdx := i
		currentValue := arr[targetedIdx]
		for j := i - 1; j >= 0; j-- {
			if arr[j] > currentValue {
				arr[j+1] = arr[j]
				targetedIdx = j
			}
		}
		if targetedIdx != i {
			arr[targetedIdx] = currentValue
		}
	}
	return arr
}
