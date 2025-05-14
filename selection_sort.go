package main

func selectionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {

		smallIdx := i

		for j := (i + 1); j < len(arr); j++ {
			if arr[j] < arr[smallIdx] {
				smallIdx = j
			}
		}

		if smallIdx != i {
			arr[i], arr[smallIdx] = arr[smallIdx], arr[i]
		}
	}

	return arr
}
