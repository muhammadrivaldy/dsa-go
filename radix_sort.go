package main

func radixSort(arr []int) []int {

	max := getMax(arr)
	exp := 1

	for (max / exp) > 0 {

		radixArr := make([][]int, 10)
		for _, i := range arr {
			r := (i / exp) % 10
			radixArr[r] = append(radixArr[r], i)
		}

		arr = []int{}
		for _, i := range radixArr {
			arr = append(arr, i...)
		}

		exp *= 10
	}

	return arr
}
