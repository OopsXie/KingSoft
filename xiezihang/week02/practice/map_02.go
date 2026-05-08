package main

import (
	"fmt"
)

func main() {
	var input string

	fmt.Println("请输入一个由字母和数字组成的字符串：")
	fmt.Scanln(&input)

	mapCount := make(map[rune]int)

	for _, char := range input {
		mapCount[char]++
	}

	fmt.Println("字符出现的次数：")
	for char, count := range mapCount {
		fmt.Printf("%c: %d\n", char, count)
	}

	var maxChar rune
	var maxCount int
	maxCount = 0
	for char, count := range mapCount {
		if count > maxCount {
			maxCount = count
			maxChar = char
		}
	}

	fmt.Printf("出现次数最多的字符是：%c，出现了 %d 次\n", maxChar, maxCount)
}
