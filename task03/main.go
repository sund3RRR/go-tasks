/*
	Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов
	(22+32+42….) с использованием конкурентных вычислений.
*/

package main

import (
	"fmt"
	"sync"
)

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

func main() {
	arr := []int{2, 4, 6, 8, 10}

	fmt.Printf("Sqr sum: %d\n", sqrSum(&arr))
}
