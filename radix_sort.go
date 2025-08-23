package main

func radixSort(arr []int) []int {
	max := getMax(arr)
	exp := 1

	for (max / exp) > 0 {
		radixArr := make([][]int, 10)
		for _, val := range arr {
			radix := (val / exp) % 10
			radixArr[radix] = append(radixArr[radix], val)
		}

		arr = []int{}
		for _, i := range radixArr {
			arr = append(arr, i...)
		}

		exp *= 10
	}

	return arr
}
