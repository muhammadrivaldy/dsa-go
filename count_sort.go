package main

func countSort(arr []int) []int {

	groupArr := [][]int{}
	for i := 0; i <= getMax(arr); i++ {
		groupArr = append(groupArr, []int{})
	}

	for _, val := range arr {
		groupArr[val] = append(groupArr[val], val)
	}

	newArr := []int{}
	for _, arr := range groupArr {
		newArr = append(newArr, arr...)
	}

	return newArr
}
