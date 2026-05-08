package main

import (
	"fmt"
)

func PalindromicNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	runes := []rune(s) // 将字符串转换为 []rune
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("This is PalindromicNumber test")
	fmt.Println()
	fmt.Println(PalindromicNumber("12321"))
	fmt.Println(PalindromicNumber("12345"))
	fmt.Println(PalindromicNumber("abcba"))

	fmt.Println(PalindromicNumber("谢子航子谢"))
	fmt.Println(PalindromicNumber("谢子航"))

	fmt.Println(PalindromicNumber("上海自来水来自海上"))
}
