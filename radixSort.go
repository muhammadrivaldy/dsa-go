package main

func radixSort(arr []int) []int {
	max := getMax(arr)
	exp := 1

	for (max / exp) > 0 {
		radixArr := make([][]int, 10)
		for _, i := range arr {
			val := (i / exp) % 10
			radixArr[val] = append(radixArr[val], i)
		}

		arr = []int{}
		for _, i := range radixArr {
			arr = append(arr, i...)
		}

		exp *= 10
	}
	return arr
}
