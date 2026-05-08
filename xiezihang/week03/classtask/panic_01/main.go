package main

import (
	"fmt"
)

func divide(a, b int) float64 {
	if b == 0 {
		panic("除数不能为零")
	}
	return float64(a) / float64(b)
}

func main() {
	a := 10
	b := 2
	var result float64

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("错误:", r)
		}
	}()

	result = divide(a, b)
	fmt.Println("结果:", result)

	result = divide(a, 0)
	fmt.Println("结果:", result)

	result = divide(a, 10)
	fmt.Println("结果:", result) //不打印

}
