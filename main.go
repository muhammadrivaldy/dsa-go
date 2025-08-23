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
	fmt.Println("7. Radix Sort")
	fmt.Println("8. Linear Search")
	fmt.Println("9. Binary Search")
	fmt.Println("10. Singly Linked List")
	fmt.Println("11. Doubly Linked List")
	fmt.Println("12. Circular Linked List")
	fmt.Print("Enter the number of the algorithm you want to run: ")
	var input int
	fmt.Scan(&input)

	switch input {
	case 1:
		fmt.Println("Merge Sort")
		nums := []int{9, 10, 4, 8, 3, 5, 2, 7, 6, 1}
		fmt.Println("Before sorting:", nums)
		fmt.Println("After sorting:", mergeSort(nums))
	case 2:
		fmt.Println("Bubble Sort")
		nums := []int{9, 10, 4, 8, 3, 5, 2, 7, 6, 1}
		fmt.Println("Before sorting:", nums)
		fmt.Println("After sorting:", bubbleSort(nums))
	case 3:
		fmt.Println("Selection Sort")
		nums := []int{9, 10, 4, 8, 3, 5, 2, 7, 6, 1}
		fmt.Println("Before sorting:", nums)
		fmt.Println("After sorting:", selectionSort(nums))
	case 4:
		fmt.Println("Insertion Sort")
		nums := []int{9, 10, 4, 8, 3, 5, 2, 7, 6, 1}
		fmt.Println("Before sorting:", nums)
		fmt.Println("After sorting:", insertionSort(nums))
	case 5:
		fmt.Println("Quick Sort")
		nums := []int{9, 10, 4, 8, 3, 5, 2, 7, 6, 1}
		fmt.Println("Before sorting:", nums)
		fmt.Println("After sorting:", quickSort(nums, 0, len(nums)-1))
	case 6:
		fmt.Println("Count Sort")
		nums := []int{4, 2, 5, 8, 7, 7, 5, 5, 3, 1, 8, 7, 6, 8, 6, 7}
		fmt.Println("Before sorting:", nums)
		fmt.Println("After sorting:", countSort(nums))
	case 7:
		fmt.Println("Radix Sort")
		nums := []int{32, 128, 20, 2, 881, 29, 95, 77}
		fmt.Println("Before sorting:", nums)
		fmt.Println("After sorting:", radixSort(nums))
	case 8:
		fmt.Println("Linear Search")
		nums := []int{9, 10, 4, 8, 3, 5, 2, 7, 6, 1}
		fmt.Printf("Array value: %v, Search value: %d\n", nums, 7)
		fmt.Println("Index:", linearSearch(nums, 7))
	case 9:
		fmt.Println("Binary Search")
		nums := []int{2, 20, 29, 32, 77, 95, 128, 881}
		fmt.Printf("Array value: %v, Search value: %d\n", nums, 29)
		fmt.Println("Index:", binarySearch(nums, 29))
	case 10:
		fmt.Println("Singly Linked List")
		nums := []int{2, 20, 29, 32, 77, 95, 128, 881}
		resp := singlyLinkedList(nums)
		singlyPrint(resp)
	case 11:
		fmt.Println("Doubly linked list")
		nums := []int{2, 20, 29, 32, 77, 95, 128, 881}
		resp := doublyLinkedList(nums)
		doublyPrint(resp)
	case 12:
		fmt.Println("Circular Linked List")
		nums := []int{2, 20, 29, 32, 77, 95, 128, 881}
		resp := circularLinkedList(nums)
		circularPrint(resp)
	default:
		fmt.Println("Invalid input")
	}
}
