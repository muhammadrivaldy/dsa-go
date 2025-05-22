package main

func insertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minVal := arr[i]
		targetedIdx := i
		for j := i - 1; j >= 0; j-- {
			if arr[j] > minVal {
				arr[j+1] = arr[j]
				targetedIdx = j
			}
		}
		arr[targetedIdx] = minVal
	}

	return arr
}
