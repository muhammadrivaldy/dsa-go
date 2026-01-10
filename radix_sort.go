package main

func radixSort(arr []int) []int {

	max := getMax(arr)
	exp := 1

	for (max / exp) > 0 {
		rArr := make([][]int, 10)
		for _, i := range arr {
			rVal := (i / exp) % 10
			rArr[rVal] = append(rArr[rVal], i)
		}

		arr = []int{}
		for _, i := range rArr {
			arr = append(arr, i...)
		}

		exp *= 10
	}

	return arr
}
