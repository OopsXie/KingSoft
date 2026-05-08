package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is operator test")
	var (
		a = 6
		b = 10
	)
	fmt.Printf("a + b = %d\n", a+b)
	fmt.Printf("a - b = %d\n", a-b)
	fmt.Printf("a * b = %d\n", a*b)
	fmt.Printf("a / b = %d\n", a/b)  //Printf和Println是有区别的
	fmt.Printf("b %% a = %d\n", b%a) // %d是占位符，%%是转义符

	fmt.Println()

	fmt.Printf("%d == %d is %t\n", a, b, a == b)
	fmt.Printf("%d != %d is %t\n", a, b, a != b)
	fmt.Printf("%d > %d is %t\n", a, b, a > b)
	fmt.Printf("%d < %d is %t\n", a, b, a < b)
	fmt.Printf("%d >= %d is %t\n", a, b, a >= b)
	fmt.Printf("%d <= %d is %t\n", a, b, a <= b)
}
