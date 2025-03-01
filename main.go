package main

import (
	"fmt"
)

func main() {

	fmt.Println("These are the list of algorithms that I have implemented so far:")
	fmt.Println("1. Merge Sort")
	fmt.Print("Enter the number of the algorithm you want to run: ")
	var input int
	fmt.Scan(&input)

	switch input {
	case 1:
		fmt.Println("Merge Sort")
		nums := []int{9, 10, 4, 8, 3, 5, 2, 7, 6, 1}
		fmt.Println("Before sorting: ", nums)
		fmt.Println("After sorting: ", mergeSort(nums))
	default:
		fmt.Println("Invalid input")
	}
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
		} else {
			break
		}
	}

	newArr = append(newArr, leftArr[leftIdx:]...)
	newArr = append(newArr, rightArr[rightIdx:]...)

	return newArr
}
