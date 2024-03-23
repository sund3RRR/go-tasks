package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
)

// Simple worker which just print data from channel
func worker(c chan string, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()

	// Read data from channel and break loop on channel closed
	for str := range c {
		fmt.Printf("Worker %d:\n\tData received from channel: %s\n", workerID, str)
	}

	// Marker, which indicates, that worker is finished successfully
	fmt.Printf("Worker %d is finished\n", workerID)
}

func doTask() {
	fmt.Println("Enter number of workers")
	// Read number of workers from cmd in string
	var workers_str string
	_, err := fmt.Scanln(&workers_str)
	if err != nil {
		log.Fatal(err)
	}

	// Parse workers count using strconv
	workers, err := strconv.ParseInt(workers_str, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	// Create main data channel
	c := make(chan string)

	// Create WaitGroup for workers
	wg := sync.WaitGroup{}

	for i := 0; i < int(workers); i++ {
		wg.Add(1)
		// Start worker in separate goroutine
		go worker(c, &wg, i)
	}

	// Create os channel and subscribe on os.Interrupt
	os_c := make(chan os.Signal, 1)
	signal.Notify(os_c, os.Interrupt)

	// Create bufio reader
	reader := bufio.NewReader(os.Stdin)

loop:
	for {
		select {
		case <-os_c:
			// Close channel on os.Interrupt signal.
			// The main idea is that all workers will stop working
			// on closed channel because 'for' loop, which reads dataunexpectedly
			// from channel will break.
			close(c)
			// Break reader loop for main goroutine.
			// If you use return instead of break, the main goroutine
			// will finish executing the process as soon as it finishes its business.
			// This means that the workers will not finish executing themselves
			// and will be terminated unexpectedly.
			break loop
		default:
		}
		data, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		// Send data to channel
		c <- data
	}

	// Wait until all workers will successfully finish their job
	wg.Wait()

}
func main() {
	doTask()
}
