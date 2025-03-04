package main

func bubbleSort(arr []int) []int {

	for {

		swapped := false

		for i := 0; i < (len(arr) - 1); i++ {

			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return arr
}
