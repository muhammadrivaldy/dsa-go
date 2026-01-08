package main

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	lArr := mergeSort(arr[:mid])
	rArr := mergeSort(arr[mid:])
	lIdx, rIdx := 0, 0
	newArr := []int{}

	for len(lArr) > lIdx && len(rArr) > rIdx {
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
