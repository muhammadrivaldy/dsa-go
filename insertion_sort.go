package main

func insertionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {

		currentVal := arr[i]
		targetedIdx := i

		for j := i; j > 0; j-- {
			if arr[j-1] > currentVal {
				arr[j] = arr[j-1]
				targetedIdx--
			} else {
				break
			}
		}

		arr[targetedIdx] = currentVal
	}

	return arr
}
