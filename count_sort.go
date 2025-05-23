package main

func countSort(arr []int) []int {
	max := getMax(arr)
	countArr := make([][]int, max+1)
	for len(arr) > 0 {
		countArr[arr[0]] = append(countArr[arr[0]], arr[0])
		arr = arr[1:]
	}

	for _, val := range countArr {
		arr = append(arr, val...)
	}
	return arr
}
