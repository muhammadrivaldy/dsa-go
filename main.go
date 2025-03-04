package main

import (
	"fmt"
)

func main() {

	fmt.Println("These are the list of algorithms that I have implemented so far:")
	fmt.Println("1. Merge Sort")
	fmt.Println("2. Bubble Sort")
	fmt.Println("3. Selection Sort")
	fmt.Println("4. Insertion Sort")
	fmt.Println("5. Quick Sort")
	fmt.Println("6. Count Sort")
	fmt.Print("Enter the number of the algorithm you want to run: ")
	var input int
	fmt.Scan(&input)

	switch input {
	case 1:
		fmt.Println("Merge Sort")
		nums := []int{9, 10, 4, 8, 3, 5, 2, 7, 6, 1}
		fmt.Println("Before sorting: ", nums)
		fmt.Println("After sorting: ", mergeSort(nums))
	case 2:
		fmt.Println("Bubble Sort")
		nums := []int{9, 10, 4, 8, 3, 5, 2, 7, 6, 1}
		fmt.Println("Before sorting: ", nums)
		fmt.Println("After sorting: ", bubbleSort(nums))
	case 3:
		fmt.Println("Selection Sort")
		nums := []int{9, 10, 4, 8, 3, 5, 2, 7, 6, 1}
		fmt.Println("Before sorting: ", nums)
		fmt.Println("After sorting: ", selectionSort(nums))
	case 4:
		fmt.Println("Insertion Sort")
		nums := []int{9, 10, 4, 8, 3, 5, 2, 7, 6, 1}
		fmt.Println("Before sorting: ", nums)
		fmt.Println("After sorting: ", insertionSort(nums))
	case 5:
		fmt.Println("Quick Sort")
		nums := []int{9, 10, 4, 8, 3, 5, 2, 7, 6, 1}
		fmt.Println("Before sorting: ", nums)
		fmt.Println("After sorting: ", quickSort(nums, 0, len(nums)-1))
	case 6:
		fmt.Println("Count Sort")
		nums := []int{4, 2, 5, 8, 7, 7, 5, 5, 3, 1, 8, 7, 6, 8, 6, 7}
		fmt.Println("Before sorting: ", nums)
		fmt.Println("After sorting: ", countSort(nums))
	default:
		fmt.Println("Invalid input")
	}
}
