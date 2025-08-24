package main

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		selection := i
		for j := (i + 1); j < len(arr); j++ {
			if arr[selection] > arr[j] {
				selection = j
			}
		}
		arr[selection], arr[i] = arr[i], arr[selection]
	}
	return arr
}
