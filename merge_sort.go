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

	for {

		if len(leftArr) > leftIdx && len(rightArr) > rightIdx {

			if leftArr[leftIdx] > rightArr[rightIdx] {
				newArr = append(newArr, rightArr[rightIdx])
				rightIdx++
			} else {
				newArr = append(newArr, leftArr[leftIdx])
				leftIdx++
			}
		} else {
			break
		}
	}

	newArr = append(newArr, leftArr[leftIdx:]...)
	newArr = append(newArr, rightArr[rightIdx:]...)

	return newArr
}
