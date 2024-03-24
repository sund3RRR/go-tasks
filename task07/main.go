package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// Protect our map data with mutex
type Map struct {
	sync.Mutex
	data map[int]int
}

func (m *Map) Set(key int, value int) {
	// Lock mutex to set data
	m.Lock()

	m.data[key] = value

	// Unlock mutex for other workers
	m.Unlock()
}

func main() {
	const DATA_COUNT = 400
	m := Map{data: make(map[int]int)}

	wg := sync.WaitGroup{}
	wg.Add(DATA_COUNT)

	for i := 0; i < DATA_COUNT; i++ {
		// Start worker with WaitGroup and ID
		go func(workerID int) {
			defer wg.Done()
			// Set random data to map
			m.Set(rand.Intn(DATA_COUNT), rand.Intn(DATA_COUNT))
			fmt.Printf("Worker %d is done\n", workerID)
		}(i)
	}
	wg.Wait()
	fmt.Println(m.data)
}
