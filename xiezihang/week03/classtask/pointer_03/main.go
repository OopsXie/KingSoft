package main

import (
	"fmt"
)

func doubleValues(arr *[5]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] *= 2
	}
}

func main() {
	var nums = [5]int{1, 2, 3, 4, 5}
	fmt.Println("nums数组：", nums)
	doubleValues(&nums)
	fmt.Println("翻倍后：", nums)
}
