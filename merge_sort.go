package main

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	mid := len(arr) / 2
	leftIdx := 0
	rightIdx := 0
	newArr := []int{}

	leftArr := mergeSort(arr[:mid])
	rightArr := mergeSort(arr[mid:])

	for len(leftArr) > leftIdx && len(rightArr) > rightIdx {
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
