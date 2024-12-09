package main

import "fmt"

func main() {
	var arrToBeSearched = []int{2, 4, 5, 6, 7, 2, 4, 5, 1, 6, 8, 76}
	target := 1
	result := linearSearch(arrToBeSearched, target)
	if result != -1 {
		fmt.Println("target found at index", result)
	} else {
		fmt.Println("target not found!")
	}
}
func linearSearch(myArray []int, target int) (result int) {
	for i := 0; i < len(myArray); i++ {
		if myArray[i] == target {
			return i
		}
	}
	return -1
}
