package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func searchFile(filePath, keyword string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), keyword) {
			results <- fmt.Sprintf("文件: %s, 行号: %d", filePath, lineNumber)
		}
		lineNumber++
	}
}

func searchDirectory(dir, keyword string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !info.IsDir() {
			wg.Add(1)
			go searchFile(path, keyword, results, wg)
		}
		return nil
	})
	if err != nil {
		fmt.Println("遍历目录时出错:", err)
	}
}

func main() {
	dir := "F:/GoProject/git.wpsit/xiezihang/week03" // 设置搜索目录
	keyword := "main"                                // 设置搜索关键字

	results := make(chan string)
	var wg sync.WaitGroup

	wg.Add(1)
	go searchDirectory(dir, keyword, results, &wg)

	go func() {
		wg.Wait()
		close(results)
	}()

	matchedFiles := make(map[string]bool)
	for result := range results {
		fmt.Println(result)
		filePath := strings.Split(result, ",")[0][4:] // 提取文件路径
		matchedFiles[filePath] = true
	}

	fmt.Printf("匹配的文件总数: %d\n", len(matchedFiles))
}
