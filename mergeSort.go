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
		lVal := leftArr[leftIdx]
		rVal := rightArr[rightIdx]

		if lVal < rVal {
			newArr = append(newArr, lVal)
			leftIdx++
		} else {
			newArr = append(newArr, rVal)
			rightIdx++
		}
	}

	newArr = append(newArr, leftArr[leftIdx:]...)
	newArr = append(newArr, rightArr[rightIdx:]...)

	return newArr
}
