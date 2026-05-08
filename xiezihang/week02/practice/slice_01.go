package main

import (
	"fmt"
)

func main() {
	//切片索引从0开始，选取第三个到第七个元素，所以即选取索引2-6的元素

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	subslice := slice[2:7] //[low, high) 左闭右开 3,4,5,6,7

	subslice = append(subslice, 100)

	var sum int
	sum = 0

	for _, v := range subslice {
		sum += v
		//fmt.Println(subslice, v)
	}

	fmt.Println("总和是:", sum)
}
