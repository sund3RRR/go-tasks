package main

import (
	"fmt"
	"sync"
	"time"
)

// Create counter struct
type Counter struct {
	// Protect our counter field with Mutex
	sync.Mutex
	count int
}

func (c *Counter) Increment(done chan int, workerID int, wg *sync.WaitGroup) {
	// defer wait group done signal for main
	defer wg.Done()

	// Create self counter for debug info
	selfCounter := 0
	for {
		select {
		// Handle done signal
		case <-done:
			fmt.Printf("Worker %d: done, %d ticks\n", workerID, selfCounter)
			return
		default:
			// Lock mutex to increment counter
			c.Lock()
			c.count++
			// Unlock mutex for other workers
			c.Unlock()

			// Increment debug counter
			selfCounter++
		}
	}
}

func main() {
	const (
		WORKER_COUNT = 8
		TIME         = 5
	)

	counter := Counter{}

	wg := sync.WaitGroup{}

	wg.Add(WORKER_COUNT)
	done := make(chan int)

	// Create workers
	for i := 0; i < WORKER_COUNT; i++ {
		go counter.Increment(done, i, &wg)
	}

	time.Sleep(TIME * time.Second)

	// Send end signal to workers
	done <- 0
	close(done)

	// Wait until all workers successfully done their job
	wg.Wait()

	fmt.Printf("Count: %d\n", counter.count)
}
