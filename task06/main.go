package main

import (
	"context"
	"fmt"
	"time"
)

func ContextMethod(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Cancel context received, terminating goroutine")
			return
		default:
			fmt.Println("Some work...")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
func PeriodicPollingMethod(done chan int) {
	for {
		select {
		case <-done:
			fmt.Println("Exit signal received, terminating goroutine")
			return
		default:
			fmt.Println("Some work...")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
func CloseChannelMethod(c chan string) {
	// Also can be implemented like this
	//
	// for data := range c {
	// 	fmt.Print(data)
	// }
	for {
		v, ok := <-c
		if !ok {
			fmt.Println("Channel closed, terminating goroutine")
			return
		}
		fmt.Print(v)
	}
}
func main() {
	//
	// The first method how to terminate goroutine is
	// to use the channel’s close mechanism.
	//
	c := make(chan string, 6)
	go CloseChannelMethod(c)
	c <- "Hello, "
	c <- "world!\n"
	close(c)
	time.Sleep(time.Second)

	//
	// The seconds is to use periodic polling `done` channel
	//
	done := make(chan int)
	go PeriodicPollingMethod(done)
	time.Sleep(time.Second)
	done <- 0
	close(done)
	time.Sleep(time.Second)

	//
	// The third is to use context with cancel signal
	//
	ctx, cancel := context.WithCancel(context.Background())
	go ContextMethod(ctx)
	time.Sleep(time.Second)
	cancel()
	time.Sleep(time.Second)
}
