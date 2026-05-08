package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Record struct {
	Record_GameID       int       //游戏ID
	Record_StartTime    time.Time //开始时间
	Record_EndTime      time.Time //结束时间
	Record_Duration     float64   //持续时间
	Record_TargetNumber int       //随机数字
	Record_Guesses      []int     //猜测的数字
	Record_Difficulty   string    //难度
	Record_Result       string    //结果
}

func RandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) + 1
}

func GuessNumber(number int, a int, guesses *[]int) bool {
	var guess int //猜测的数字
	var flag = false
	for i := 0; i < a; i++ {
		fmt.Printf("第%d次猜测，请输入您的数字（1-100）:", i+1)
		_, err := fmt.Scanln(&guess)

		if err != nil {
			fmt.Println("输入错误，请输入数字")
			var dump string
			fmt.Scanln(&dump)
			i-- //不算这次机会
			continue
		}

		if guess < 1 || guess > 100 {
			fmt.Println("输入错误，请输入1-100之间的数字")
			i-- //不算这次机会
			continue
		}

		*guesses = append(*guesses, guess)
		if guess == number {
			flag = true
			fmt.Printf("恭喜您猜对了!您在第%d次猜测中成功。", i+1)
			break
		} else if guess > number {
			fmt.Println("您猜的数字大了")
		} else {
			fmt.Println("您猜的数字小了")
		}
	}
	if !flag {
		fmt.Println("很遗憾，您的机会已用完，正确答案是：", number)
	}
	return flag
}

func exit(flag bool) bool {
	if !flag {
		fmt.Println("给定的次数您没猜对，游戏结束，您可重新开始游戏或者退出。是否继续游戏？(Y/N)")
		var s string
		fmt.Scanln(&s)
		if s == "Y" || s == "y" {
			return true
		} else {
			return false
		}
	} else {
		fmt.Println("是否继续游戏？(Y/N)")
		var s string
		fmt.Scanln(&s)
		if s == "Y" || s == "y" {
			return true
		} else {
			return false
		}
	}
}

func WriteRecord(record Record) {
	file, err := os.OpenFile("game.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("文件打开失败:", err)
		return
	}
	defer file.Close()

	line := fmt.Sprintf("Game %d | 难度: %s | 开始: %s | 结束: %s | 用时: %.2fs | 随机数字: %d | 猜测: %v | 结果: %s\n",
		record.Record_GameID,
		record.Record_Difficulty,
		record.Record_StartTime.Format("2006-01-02 15:04:05"),
		record.Record_EndTime.Format("2006-01-02 15:04:05"),
		record.Record_Duration,
		record.Record_TargetNumber,
		record.Record_Guesses,
		record.Record_Result,
	)
	file.WriteString(line)
}

func OutPrint() {
	fmt.Println("欢迎来到猜数字游戏！")
	fmt.Println("游戏规则如下：")
	fmt.Println("1. 计算机将在1到100 之间随机选择一个数字。")
	fmt.Println("2. 您可以选择难度级别(简单、中等、困难)，不同难度对应不同的猜测机会")
	fmt.Println("3. 请输入您的猜测。")
	fmt.Println("")
	fmt.Println("请选择难度级别(简单/中等/困难)：")
	fmt.Println("0.退出游戏")
	fmt.Println("1.简单(10 次机会)")
	fmt.Println("2.中等(5 次机会)")
	fmt.Println("3.困难(3 次机会)")
}

func main() {
	GameID := 1
	for {
		record_1 := Record{}
		record_1.Record_GameID = GameID
		OutPrint()
		var level int  //难度
		var number int //随机数
		var flag bool  //是否猜对
		var tmp bool   //是否继续游戏

		var end time.Time
		var duration float64

		number = RandomNumber()
		fmt.Print("输入选择:")
		fmt.Scanln(&level)
		start := time.Now() //获取当前时间
		fmt.Println("开始游戏")
		if level == 0 {
			fmt.Println("游戏结束")
			break
		} else if level == 1 {
			flag = GuessNumber(number, 10, &record_1.Record_Guesses)
			end = time.Now()
			duration = end.Sub(start).Seconds() //计算时间差
			fmt.Println("游戏结束，总共用时：", duration)
		} else if level == 2 {
			flag = GuessNumber(number, 5, &record_1.Record_Guesses)
			end = time.Now()
			duration = end.Sub(start).Seconds() //计算时间差
			fmt.Println("游戏结束，总共用时：", duration)
		} else if level == 3 {
			flag = GuessNumber(number, 3, &record_1.Record_Guesses)
			end = time.Now()
			duration = end.Sub(start).Seconds() //计算时间差
			fmt.Println("游戏结束，总共用时：", duration)
		} else {
			fmt.Println("输入有误，请重新输入")
		}

		record_1.Record_StartTime = start
		record_1.Record_EndTime = end
		record_1.Record_Duration = duration
		record_1.Record_TargetNumber = number

		switch level {
		case 1:
			record_1.Record_Difficulty = "简单"
		case 2:
			record_1.Record_Difficulty = "中等"
		case 3:
			record_1.Record_Difficulty = "困难"
		}

		if flag {
			record_1.Record_Result = "成功"
		} else {
			record_1.Record_Result = "失败"
		}

		WriteRecord(record_1)
		GameID++
		tmp = exit(flag)
		if tmp {
			continue
		} else {
			break
		}
	}
}
