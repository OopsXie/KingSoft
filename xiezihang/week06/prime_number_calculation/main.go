package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func findPrimes(start, end int, wg *sync.WaitGroup, ch chan []int) {
	defer wg.Done()
	primes := []int{}
	for i := start; i <= end; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	ch <- primes
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("用法: go run main.go [起始值] [结束值]")
		return
	}

	start, err1 := strconv.Atoi(os.Args[1])
	end, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil {
		fmt.Println("参数必须是整数")
		return
	}
	if start < 0 || end < 0 {
		fmt.Println("起始值和结束值必须是非负整数")
		return
	}
	if start > end {
		fmt.Println("起始值必须小于等于结束值")
		return
	}

	startTime := time.Now()

	segment := (end - start + 1) / 4
	var wg sync.WaitGroup
	ch := make(chan []int, 4)

	for i := 0; i < 4; i++ {
		s := start + i*segment
		e := s + segment - 1
		if i == 3 {
			e = end
		}
		wg.Add(1)
		go findPrimes(s, e, &wg, ch)
	}

	wg.Wait()
	close(ch)

	// 合并所有结果
	allPrimes := []int{}
	for part := range ch {
		allPrimes = append(allPrimes, part...)
	}

	//sort.Ints(allPrimes)排序

	filename := fmt.Sprintf("primes_%d_%d.txt", start, end)
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("文件创建失败: %v\n", err)
		return
	}
	defer file.Close()

	for _, prime := range allPrimes {
		fmt.Fprintln(file, prime)
	}

	elapsed := time.Since(startTime)

	fmt.Printf("找到 %d 个素数\n", len(allPrimes))
	fmt.Printf("运行时间: %s\n", elapsed)
	fmt.Printf("结果已保存到文件: %s\n", filename)
}
