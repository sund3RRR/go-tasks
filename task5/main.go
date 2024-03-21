package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func readerWorker(c chan string, ctx context.Context) {
	for str := range c {
		select {
		case <-ctx.Done():
			fmt.Println("Reader worker is finished")
			return
		default:
			fmt.Printf("Data received from channel: %s", str)
		}
	}
	fmt.Println("Reader worker is finished")
}

func writerWorker(c chan string, ctx context.Context) {
	reader := bufio.NewReader(os.Stdin)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Writer worker is finished")
			return
		default:
			data, err := reader.ReadString('\n')
			if err != nil {
				log.Printf("ERROR: %s", err)
				break
			}
			c <- data
		}
	}
	fmt.Println("Writer worker is finished")
}

func main() {
	var seconds_str string

	_, err := fmt.Scanln(&seconds_str)
	if err != nil {
		log.Fatal(err)
	}

	seconds, err := strconv.ParseInt(seconds_str, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan string, 1)

	ctx, cancel := context.WithCancel(context.Background())

	go readerWorker(c, ctx)
	go writerWorker(c, ctx)

	time.Sleep(time.Second * time.Duration(seconds))
	close(c)
	cancel()
	time.Sleep(time.Second)
}
