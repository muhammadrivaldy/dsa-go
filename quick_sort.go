package main

func quickSort(arr []int, low, high int) []int {

	if low <= high {
		pivot := quickSortPivot(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}

	return arr
}

func quickSortPivot(arr []int, low, high int) int {

	i := low

	for j := low; j < high; j++ {
		if arr[j] <= arr[high] {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]

	return i
}
