package main

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	leftArr := mergeSort(arr[:mid])
	rightArr := mergeSort(arr[mid:])
	leftIdx, rightIdx := 0, 0
	arr = []int{}

	for len(leftArr) > leftIdx && len(rightArr) > rightIdx {
		if leftArr[leftIdx] < rightArr[rightIdx] {
			arr = append(arr, leftArr[leftIdx])
			leftIdx++
		} else {
			arr = append(arr, rightArr[rightIdx])
			rightIdx++
		}
	}

	arr = append(arr, leftArr[leftIdx:]...)
	arr = append(arr, rightArr[rightIdx:]...)
	return arr
}
