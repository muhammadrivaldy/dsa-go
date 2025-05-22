package main

func quickSort(arr []int, low, high int) []int {
	if low < high {
		pivot := quickSortPivot(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}

	return arr
}

func quickSortPivot(arr []int, low, high int) int {
	pivot := low
	for i := low; i < high; i++ {
		if arr[i] < arr[high] {
			arr[i], arr[pivot] = arr[pivot], arr[i]
			pivot++
		}
	}

	arr[pivot], arr[high] = arr[high], arr[pivot]

	return pivot
}
