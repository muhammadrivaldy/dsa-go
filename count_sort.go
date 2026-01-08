package main

func countSort(arr []int) []int {

	max := getMax(arr)
	countArr := make([][]int, (max + 1))
	for _, i := range arr {
		countArr[i] = append(countArr[i], i)
	}

	newArr := []int{}
	for _, i := range countArr {
		if len(i) > 0 {
			newArr = append(newArr, i...)
		}
	}

	return newArr
}
