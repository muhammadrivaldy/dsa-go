package main

func radixSort(arr []int) []int {
	max := getMax(arr)
	exp := 1
	for (max / exp) > 0 {
		radixArr := make([][]int, 10)
		for _, val := range arr {
			idxVal := (val / exp) % 10
			radixArr[idxVal] = append(radixArr[idxVal], val)
		}

		arr = []int{}
		for _, val := range radixArr {
			arr = append(arr, val...)
		}
		exp *= 10
	}
	return arr
}
