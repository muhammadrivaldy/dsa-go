package main

func bubbleSort(arr []int) []int {
	for {
		isSwap := false
		for i := 1; i < len(arr); i++ {
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				isSwap = true
			}
		}
		if !isSwap {
			break
		}
	}
	return arr
}
