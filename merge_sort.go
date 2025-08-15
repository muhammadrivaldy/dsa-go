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

	for len(lArr) > lIdx && len(rArr) > rIdx {

		lVal := lArr[lIdx]
		rVal := rArr[rIdx]

		if lVal < rVal {
			newArr = append(newArr, lVal)
			lIdx++
		} else {
			newArr = append(newArr, rVal)
			rIdx++
		}
	}

	newArr = append(newArr, lArr[lIdx:]...)
	newArr = append(newArr, rArr[rIdx:]...)

	return newArr
}
