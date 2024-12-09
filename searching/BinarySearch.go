package main

import "fmt"

func main() {
	myArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7

	result := binarySearch(myArray, target)
	if result != -1 {
		fmt.Println("target found at index", result)
	} else {
		fmt.Println("target not found,", result)
	}
}
func binarySearch(myArray []int, target int) (result int) {
	left := 0
	right := len(myArray) - 1
	for left < right {
		mid := (left + right) / 2
		if myArray[mid] == target {
			return mid
		} else if left < mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
