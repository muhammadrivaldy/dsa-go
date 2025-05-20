package main

func countSort(arr []int) []int {

	max := getMax(arr)
	countedArr := make([][]int, max+1)

	for _, i := range arr {
		countedArr[i] = append(countedArr[i], i)
	}

	newArr := []int{}
	for _, i := range countedArr {
		newArr = append(newArr, i...)
	}

	return newArr
}
