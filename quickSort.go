package main

func quickSort(arr []int, min, max int) []int {

	if min < max {
		pivot := quickSortPivot(arr, min, max)
		quickSort(arr, 0, (pivot - 1))
		quickSort(arr, (pivot + 1), max)
	}

	return arr
}

func quickSortPivot(arr []int, min, max int) int {

	pivot := min

	for i := min; i < len(arr); i++ {
		if arr[i] < arr[max] {
			arr[i], arr[pivot] = arr[pivot], arr[i]
			pivot++
		}
	}

	arr[max], arr[pivot] = arr[pivot], arr[max]
	return pivot
}
