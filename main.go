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

func bubbleSort(arr []int) []int {

	for {

		swapped := false

		for i := 0; i < (len(arr) - 1); i++ {

			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return arr
}

func selectionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {

		minIndex := i

		for j := i + 1; j < len(arr); j++ {

			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}

		if minIndex != i {
			arr[minIndex], arr[i] = arr[i], arr[minIndex]
		}
	}

	return arr
}

func insertionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {

		insertIndex := i
		currentValue := arr[i]

		for j := (i - 1); j >= 0; j-- {

			if arr[j] > currentValue {
				arr[(j + 1)] = arr[j]
				insertIndex = j
			} else {
				break
			}
		}

		arr[insertIndex] = currentValue
	}

	return arr
}

func quickSort(arr []int, low, high int) []int {

	if low < high {

		pivot := quickSortPartition(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}

	return arr
}

func quickSortPartition(arr []int, low, high int) int {

	i := low

	for j := low; j < high; j++ {

		if arr[j] <= arr[high] {
			arr[i], arr[j] = arr[j], arr[i]
			i += 1
		}
	}

	arr[i], arr[high] = arr[high], arr[i]

	return i
}
