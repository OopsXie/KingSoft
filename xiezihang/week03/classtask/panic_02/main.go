package main

import (
	"fmt"
)

func accessArray(arr []int, index int) int {
	if index < 0 || index >= len(arr) {
		panic(fmt.Sprintf("数组越界"))
	}
	return arr[index]
}

func main() {
	var nums = [5]int{1, 2, 3, 4, 5}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("错误", r)
		}
	}()

	fmt.Println("索引为2的元素：", accessArray(nums[:], 2))
	fmt.Println("索引为0的元素：", accessArray(nums[:], 0))
	fmt.Println("索引为-1的元素：", accessArray(nums[:], -1))
}
