package main

func radixSort(arr []int) []int {

	max := getMax(arr)
	exp := 1

	for (max / exp) > 0 {

		radixArr := make([][]int, 10)

		for len(arr) > 0 {
			val := arr[len(arr)-1]
			arr = arr[:len(arr)-1]
			radixIdr := (val / exp) % 10
			radixArr[radixIdr] = append(radixArr[radixIdr], val)
		}

		for _, bucket := range radixArr {
			for len(bucket) > 0 {
				val := bucket[len(bucket)-1]
				bucket = bucket[:len(bucket)-1]
				arr = append(arr, val)
			}
		}

		exp *= 10
	}

	return arr
}
