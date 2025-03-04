package main

func countSort(arr []int) []int {

	groupArr := make([][]int, (getMax(arr) + 1))
	for _, val := range arr {
		groupArr[val] = append(groupArr[val], val)
	}

	newArr := []int{}
	for _, arr := range groupArr {
		newArr = append(newArr, arr...)
	}

	return newArr
}
