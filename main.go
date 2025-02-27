package main

import "fmt"

func main() {
	fmt.Println(mergeSort([]int{1, 4, 3, 2, 7}))
}

func mergeSort(arr []int) []int {

	if len(arr) == 1 {
		return arr
	}

	midArr := len(arr) / 2

	leftArr := mergeSort(arr[:midArr])
	rightArr := mergeSort(arr[midArr:])

	leftIdx := 0
	rightIdx := 0
	newArr := []int{}

	for {

		if len(leftArr) > leftIdx && len(rightArr) > rightIdx {

			if leftArr[leftIdx] > rightArr[rightIdx] {
				newArr = append(newArr, rightArr[rightIdx])
				rightIdx++
			} else {
				newArr = append(newArr, leftArr[leftIdx])
				leftIdx++
			}

		} else if len(leftArr) > leftIdx {
			newArr = append(newArr, leftArr[leftIdx:]...)
			break
		} else if len(rightArr) > rightIdx {
			newArr = append(newArr, rightArr[rightIdx:]...)
			break
		}
	}

	return newArr
}
