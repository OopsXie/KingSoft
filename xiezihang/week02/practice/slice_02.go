package main

import (
	"fmt"
)

func main() {
	//切片索引从0开始，选取第三个到第七个元素，所以即选取索引2-6的元素

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	subslice := slice[2:7]                           //[low, high) 左闭右开 3,4,5,6,7
	subslice = append(subslice, 11, 12, 13)          //在切片末尾添加三个新元素:11,12,13
	subslice = append(subslice[:4], subslice[5:]...) // 删除切片中的第5个元素 即 索引4的元素
	//fmt.Println(subslice)

	for v := range subslice {
		subslice[v] *= 2
	}

	fmt.Println("切片内容：", subslice)
	fmt.Println("切片容量：", cap(subslice))
	//fmt.Println("切片长度：", len(subslice))  //长度和容量不一样
}
