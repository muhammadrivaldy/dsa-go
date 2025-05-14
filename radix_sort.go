package main

func radixSort(arr []int) []int {

	max := getMax(arr)
	exp := 1

	for (max / exp) > 0 {

		radixArr := make([][]int, 10)

		for len(arr) > 0 {
			val := arr[len(arr)-1]
			arr = arr[:len(arr)-1]
			radixIdx := (val / exp) % 10
			radixArr[radixIdx] = append(radixArr[radixIdx], val)
		}

		for _, bucket := range radixArr {
			for i := len(bucket); i > 0; i-- {
				arr = append(arr, bucket[i-1])
			}
		}

		exp *= 10
	}

	return arr
}
