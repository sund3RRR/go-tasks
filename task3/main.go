package task3

import (
	"fmt"
	"sync"
)

const taskNum = 3

// sqrSum calculates the sum of sqr(num) for all nums in an integer array
func sqrSum(arr *[]int) int {
	wg := sync.WaitGroup{}
	c := make(chan int)

	for _, num := range *arr {
		wg.Add(1)
		// Start goroutine for all nums in array
		go func(num int) {
			defer wg.Done()
			c <- num * num
		}(num)
	}

	// Wait until all goroutines have completed their work
	go func() {
		wg.Wait()
		close(c)
	}()

	res := 0
	// Read data from a channel
	for num := range c {
		res += num
	}

	return res
}
func DoTask() {
	fmt.Printf("\nStart task %d\n\n", taskNum)

	arr := []int{2, 4, 6, 8, 10}

	fmt.Printf("Sqr sum: %d", sqrSum(&arr))

	fmt.Printf("\nEnd of the task %d\n", taskNum)
}
