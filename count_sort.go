package main

func countSort(arr []int) []int {
	max := getMax(arr)
	count2DArr := make([][]int, max+1)
	for _, val := range arr {
		count2DArr[val] = append(count2DArr[val], val)
	}

	newArr := []int{}
	for _, val := range count2DArr {
		newArr = append(newArr, val...)
	}

	return newArr
}
