package main

func radixSort(arr []int) []int {
	max := getMax(arr)
	exp := 1

	for (max / exp) > 0 {
		radixArr := make([][]int, 10)
		for len(arr) > 0 {
			val := arr[0]
			arr = arr[1:]
			radix := (val / exp) % 10
			radixArr[radix] = append(radixArr[radix], val)
		}

		for _, bucket := range radixArr {
			arr = append(arr, bucket...)
		}

		exp *= 10
	}

	return arr
}
