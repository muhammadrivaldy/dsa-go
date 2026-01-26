package main

func radixSort(arr []int) []int {

	max := getMax(arr)
	exp := 1

	for (max / exp) > 0 {

		rArr := make([][]int, 10)
		for _, i := range arr {
			rIdx := (i / exp) % 10
			rArr[rIdx] = append(rArr[rIdx], i)
		}

		arr = []int{}
		for _, i := range rArr {
			arr = append(arr, i...)
		}

		exp *= 10
	}

	return arr
}
