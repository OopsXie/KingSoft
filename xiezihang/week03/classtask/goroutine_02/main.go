package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func countWord(filePath string, targetWord string, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("打开文件失败: %s，错误: %v\n", filePath, err)
		ch <- 0
		return
	}
	defer file.Close()

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		count += strings.Count(line, targetWord)
	}

	ch <- count
}

func main() {

	files := []string{"1.txt", "2.txt", "3.txt"}
	targetWord := "hello"

	ch := make(chan int, len(files))

	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go countWord(file, targetWord, ch, &wg)
	}

	wg.Wait()
	close(ch)

	totalCount := 0
	for num := range ch {
		totalCount += num
	}

	fmt.Printf("所有文件中，单词 [%s] 的总出现次数是: %d\n", targetWord, totalCount)
}
