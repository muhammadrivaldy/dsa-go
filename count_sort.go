package main

func countSort(arr []int) []int {
	max := getMax(arr)
	countArr := make([][]int, max+1)
	for _, val := range arr {
		countArr[val] = append(countArr[val], val)
	}

	newArr := []int{}
	for _, val := range countArr {
		newArr = append(newArr, val...)
	}

	return newArr
}
