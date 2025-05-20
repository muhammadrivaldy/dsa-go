package main

func selectionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		min := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min]
	}

	return arr
}
