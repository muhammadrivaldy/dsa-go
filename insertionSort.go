package main

// 3, 4, 5, 1
func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		targetIdx := i
		currentVal := arr[i]
		for j := (i - 1); j >= 0; j-- {
			if arr[j] > currentVal {
				arr[j+1] = arr[j]
				targetIdx = j
			}
		}
		arr[targetIdx] = currentVal
	}
	return arr
}
