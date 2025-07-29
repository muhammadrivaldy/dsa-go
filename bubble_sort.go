package main

func bubbleSort(arr []int) []int {

	for {
		isSwapped := false
		for i := 1; i < len(arr); i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				isSwapped = true
			}
		}

		if !isSwapped {
			break
		}
	}

	return arr
}
