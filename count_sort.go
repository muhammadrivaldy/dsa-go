package main

func countSort(arr []int, max int) []int {

	groupArr := [][]int{}
	for i := 0; i <= max; i++ {
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
