// 实现一个简单的生产者 - 消费者模型。生产者生成一系列数字并发送到通道，消费者从通道接收数字并打印。生产者在生成一定数量的数字后停止，消费者在通道关闭后停止。
package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int, count int) {
	for i := 1; i <= count; i++ {
		fmt.Printf("生产者生产了: %d\n", i)
		ch <- i
		time.Sleep(time.Millisecond * 300)
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for num := range ch {
		fmt.Printf("消费者消费了: %d\n", num)
		time.Sleep(time.Millisecond * 500)
	}
	fmt.Println("消费者停止消费。")
}

func main() {
	ch := make(chan int, 10)
	go producer(ch, 10)
	consumer(ch)
}
