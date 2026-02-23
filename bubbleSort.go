package main

func bubbleSort(arr []int) []int {

	for {
		swap := false
		for i := 1; i < len(arr); i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swap = true
			}
		}
		if !swap {
			break
		}
	}

	return arr
}
