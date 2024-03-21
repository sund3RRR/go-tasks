package main

import (
	"fmt"
	"sync"
)

const taskNum = 2

// sqr calculates squares for all nums in an integer array
func sqr(arr *[]int) {
	wg := sync.WaitGroup{}

	for _, i := range *arr {
		wg.Add(1)
		// Start goroutine for all nums in array
		go func(i int) {
			defer wg.Done()
			fmt.Println(i * i)
		}(i)
	}

	wg.Wait()
}
func main() {
	fmt.Printf("\nStart task %d\n\n", taskNum)

	arr := []int{2, 4, 6, 8, 10}

	sqr(&arr)

	fmt.Printf("\nEnd of the task %d\n", taskNum)
}
