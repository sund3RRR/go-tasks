/*
	Написать программу, которая конкурентно рассчитает значение квадратов
	чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}

	wg := sync.WaitGroup{}

	wg.Add(len(arr))

	for _, i := range arr {
		// Start goroutine for all nums in array
		go func(i int) {
			defer wg.Done()
			fmt.Println(i * i)
		}(i)
	}

	wg.Wait()
}
