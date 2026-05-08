package main

import (
	"fmt"
	"sync"
)

func calcPartSum(nums []int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0
	for _, num := range nums {
		sum += num * num
	}

	ch <- sum
}

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int, 3)

	var wg sync.WaitGroup

	partSize := (len(nums) + 2) / 3
	for i := 0; i < len(nums); i += partSize {
		end := i + partSize
		if end > len(nums) {
			end = len(nums)
		}

		wg.Add(1)
		go calcPartSum(nums[i:end], ch, &wg)
	}

	wg.Wait()

	close(ch)

	total := 0
	for sum := range ch {
		total += sum
	}

	fmt.Println("切片所有元素平方和是：", total)
}
