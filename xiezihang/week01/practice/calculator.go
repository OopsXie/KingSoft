package main

import (
	"errors"
	"fmt"
)

// calculate 执行计算并返回结果
func calculate(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, errors.New("输入错误，除数不能为 0")
		}
		return num1 / num2, nil
	default:
		return 0, errors.New("输入错误，无效的运算符")
	}

	// TODO: 实现计算逻辑
	// 根据operator执行相应的运算
	// 如果operator不是有效的运算符，返回错误

	return 0, nil
}

// getInput 获取用户输入的两个数字
func getInput() (float64, float64, error) {
	var num1, num2 float64
	fmt.Print("请输入第一个数字: ")
	_, err := fmt.Scanln(&num1)
	if err != nil {
		return 0, 0, err
	}

	fmt.Print("请输入第二个数字: ")
	_, err = fmt.Scanln(&num2)
	if err != nil {
		return 0, 0, err
	}

	return num1, num2, nil
}

func main() {
	fmt.Println("欢迎使用简单计算器！")
	fmt.Println("支持的运算：+（加）, -（减）, *（乘）, /（除）")
	fmt.Println("输入 q 退出程序")

	for {
		fmt.Print("\n请输入运算符: ")
		var operator string
		fmt.Scanln(&operator)

		if operator == "q" || operator == "Q" {
			fmt.Println("感谢使用，再见！")
			break
		}

		// 获取输入数字
		num1, num2, err := getInput()
		if err != nil {
			fmt.Printf("输入错误: %v\n", err)
			continue
		}

		// 计算结果
		result, err := calculate(num1, num2, operator)
		if err != nil {
			fmt.Printf("计算错误: %v\n", err)
			continue
		}

		// 输出结果
		fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, result)
	}
}
