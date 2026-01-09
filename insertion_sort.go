package main

func insertionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		insIdx := i
		insVal := arr[i]
		for j := (i - 1); j >= 0; j-- {
			if arr[j] > insVal {
				arr[j+1] = arr[j]
				insIdx = j
			}
		}
		arr[insIdx] = insVal
	}

	return arr
}
