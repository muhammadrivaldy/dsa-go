package main

func countSort(arr []int) []int {

	max := getMax(arr)
	cArr := make([][]int, (max + 1))

	for _, i := range arr {
		cArr[i] = append(cArr[i], i)
	}

	arr = []int{}
	for _, i := range cArr {
		arr = append(arr, i...)
	}

	return arr
}
