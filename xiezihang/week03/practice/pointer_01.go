package main

import (
	"fmt"
)

func Swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	num1, num2 := 5, 10
	fmt.Printf("交换前：num1 = %d, num2 = %d\n", num1, num2)
	Swap(&num1, &num2)
	fmt.Printf("交换后：num1 = %d, num2 = %d\n", num1, num2)
}
