package main

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	lArr := mergeSort(arr[:mid])
	rArr := mergeSort(arr[mid:])

	newArr := []int{}
	lIdx := 0
	rIdx := 0

	for lIdx < len(lArr) && rIdx < len(rArr) {
		if lArr[lIdx] < rArr[rIdx] {
			newArr = append(newArr, lArr[lIdx])
			lIdx++
		} else {
			newArr = append(newArr, rArr[rIdx])
			rIdx++
		}
	}

	newArr = append(newArr, lArr[lIdx:]...)
	newArr = append(newArr, rArr[rIdx:]...)
	return newArr
}
