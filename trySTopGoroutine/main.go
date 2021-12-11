package main

import (
	"context"
	"fmt"
	"time"
)

// 如何关闭goroutine
func main() {
	// 1、借助channel精确关闭共routine
	//stopWithChannel()

	// 2、 结合select 定期轮询channel，
	// 使用select 监控一个done的channel，如果n秒后done中接收到数据，关闭主要的channel，退出goroutine
	//stopWithSelect()

	// 3、 结合context
	stopWithContext()

}

func stopWithChannel() {
	ch := make(chan string, 6)
	go func() {
		for {
			v, ok := <-ch
			if !ok {
				fmt.Print("结束")
				return
			}
			fmt.Print(v)

			// for range 当有数据的时候，输出
			//for v := range ch {
			//	fmt.Print(v)
			//}
		}
	}()

	ch <- "我是jisiiah"
	ch <- "可你不是"
	close(ch)
	time.Sleep(time.Second)
}

func stopWithSelect() {
	ch := make(chan string, 6)
	done := make(chan struct{})

	go func() {
		for {
			select {
			case ch <- "i am josiah":
			case <-done:
				close(ch)
				return
			}
		}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		done <- struct{}{}
	}()

	for v := range ch {
		fmt.Print(v)
	}

	fmt.Print("结束")

}

func stopWithContext() {
	ch := make(chan string, 6)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case ch <- "i am josiah":
			case <-ctx.Done():
				close(ch)
				return
			}
		}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	for v := range ch {
		fmt.Print(v)
	}

	fmt.Print("结束")
}
