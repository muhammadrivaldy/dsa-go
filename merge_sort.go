package main

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	mid := len(arr) / 2
	leftArr := mergeSort(arr[:mid])
	rightArr := mergeSort(arr[mid:])

	newArr := []int{}
	leftIdx := 0
	rightIdx := 0

	for len(leftArr) > leftIdx && len(rightArr) > rightIdx {
		leftVal := leftArr[leftIdx]
		rightVal := rightArr[rightIdx]
		if leftVal < rightVal {
			newArr = append(newArr, leftVal)
			leftIdx++
		} else {
			newArr = append(newArr, rightVal)
			rightIdx++
		}
	}

	newArr = append(newArr, leftArr[leftIdx:]...)
	newArr = append(newArr, rightArr[rightIdx:]...)

	return newArr
}
