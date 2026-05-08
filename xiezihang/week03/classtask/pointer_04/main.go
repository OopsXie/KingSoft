package main

import (
	"fmt"
)

func divide(a *int, b *int, result *float64) error {
	if *b == 0 {
		return fmt.Errorf("除数不能为零")
	}
	*result = float64(*a) / float64(*b)
	return nil
}

func main() {
	var a, b int
	var result float64
	var err error

	a = 10
	b = 5
	err = divide(&a, &b, &result)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Printf("%d / %d = %.2f\n", a, b, result)
	}

	a = 10
	b = 0
	err = divide(&a, &b, &result)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Printf("%d / %d = %.2f\n", a, b, result)
	}

}
