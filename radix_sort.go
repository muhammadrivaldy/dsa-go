package main

func radixSort(arr []int) []int {
	max := getMax(arr)
	exp := 1

	for (max / exp) > 0 {
		radixArr := make([][]int, 10)
		for len(arr) > 0 {
			val := arr[0]
			arr = arr[1:]
			radixVal := (val / exp) % 10
			radixArr[radixVal] = append(radixArr[radixVal], val)
		}

		for _, i := range radixArr {
			arr = append(arr, i...)
		}

		exp *= 10
	}

	return arr
}
