package main

func bubbleSort(arr []int) []int {

	for {
		swapped := false
		for j := 1; j < len(arr); j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return arr
}
