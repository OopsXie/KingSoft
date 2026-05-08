package main

import (
	"fmt"
)

func removeDuplicates(combinedSlice []int) []int {
	uniqueMap := make(map[int]bool)
	var uniqueSlice []int

	for _, v := range combinedSlice {
		if _, exists := uniqueMap[v]; !exists {
			uniqueMap[v] = true
			uniqueSlice = append(uniqueSlice, v)
		}
	}

	return uniqueSlice
}

func main() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{3, 4, 5, 6}

	combinedSlice := append(slice1, slice2...)
	var uniqueSlice []int
	uniqueSlice = removeDuplicates(combinedSlice)
	fmt.Println("去重后的切片:", uniqueSlice)
}
