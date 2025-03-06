package main

func mergeSort(arr []int) []int {

	if len(arr) == 1 {
		return arr
	}

	midArr := len(arr) / 2

	leftArr := mergeSort(arr[:midArr])
	rightArr := mergeSort(arr[midArr:])

	leftIdx := 0
	rightIdx := 0
	newArr := []int{}

	for leftIdx < len(leftArr) && rightIdx < len(rightArr) {
		if leftArr[leftIdx] < rightArr[rightIdx] {
			newArr = append(newArr, leftArr[leftIdx])
			leftIdx++
		} else {
			newArr = append(newArr, rightArr[rightIdx])
			rightIdx++
		}
	}

	newArr = append(newArr, leftArr[leftIdx:]...)
	newArr = append(newArr, rightArr[rightIdx:]...)

	return newArr
}
