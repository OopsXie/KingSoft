package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Println("请输入字符串：")
	fmt.Scan(&str)
	fmt.Println("字符数量:", len([]rune(str)))
}
