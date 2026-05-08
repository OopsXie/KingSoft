package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	retryCount = 3
)

func downloadImage(url string, retry int) error {
	resp, err := http.Get(url)
	if err != nil {
		if retry > 0 {
			time.Sleep(2 * time.Second) // 等待后重试
			return downloadImage(url, retry-1)
		}
		return fmt.Errorf("failed to download %s after retries: %v", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("non-200 status code for %s: %d", url, resp.StatusCode)
	}

	fileName := filepath.Base(url)
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %v", fileName, err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save image %s: %v", fileName, err)
	}

	fmt.Printf("Downloaded: %s\n", fileName)
	return nil
}

func main() {
	file, err := os.Open("urls.txt")
	if err != nil {
		fmt.Printf("Failed to open urls.txt: %v\n", err)
		return
	}
	defer file.Close()

	var wg sync.WaitGroup
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			if err := downloadImage(url, retryCount); err != nil {
				fmt.Printf("Error downloading %s: %v\n", url, err)
			}
		}(url)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading urls.txt: %v\n", err)
	}

	wg.Wait()
	fmt.Println("All downloads completed.")
}
