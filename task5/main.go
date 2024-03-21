package main

import (
	"bufio"

	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

const taskNum = 5

// Simple worker which just print data from channel
func readerWorker(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Read data from channel and break loop on channel closed
	for str := range c {
		fmt.Printf("Data received from channel: %s", str)
	}
	// Marker, which indicates, that worker is finished successfully
	log.Println("INFO: Reader worker is finished")
}

func doTask() {
	var seconds_str string

	fmt.Println("Enter program TTL in seconds")
	_, err := fmt.Scanln(&seconds_str)
	if err != nil {
		log.Fatal(err)
	}

	seconds, err := strconv.ParseInt(seconds_str, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan string, 1)

	// Create wait group for correctly terminating goroutines
	wg := sync.WaitGroup{}

	wg.Add(1)
	go readerWorker(c, &wg)

	end := time.Now().Add(time.Second * time.Duration(seconds))
	reader := bufio.NewReader(os.Stdin)

	// Loop for program TTL
	for now := time.Now(); now.Unix() < end.Unix(); now = time.Now() {
		data, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("ERROR: %s", err)
			break
		}
		// Send data to channel
		c <- data
	}
	close(c)

	// Wait until reader worker will successfully finish his work
	wg.Wait()

	log.Println("INFO: Exiting program due to expiration of TTL")
}

func main() {
	fmt.Printf("\nStart task %d\n\n", taskNum)

	doTask()

	fmt.Printf("\nEnd of the task %d\n", taskNum)
}
