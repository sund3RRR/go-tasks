/*
	Разработать конвейер чисел. Даны два канала: в первый пишутся
	числа (x) из массива, во второй — результат операции x*2, после
	чего данные из второго канала должны выводиться в stdout.
*/

package main

import "fmt"

func main() {
	nums := []int{12, 545, 76, 163, 6752, 326, 6283}

	// Channel for nums from array
	numsChan := make(chan int, len(nums))

	// Write nums to channel
	for _, num := range nums {
		numsChan <- num
	}
	// Close channel because we don't need to write here anymore
	close(numsChan)

	// Channel for doubled nums
	doubledNumsChan := make(chan int, len(nums))

	// Write doubled nums to channel
	for num := range numsChan {
		doubledNumsChan <- num * 2
	}

	// Close channel because we don't need to write here anymore
	close(doubledNumsChan)

	// Read and print nums from doubled channel
	for num := range doubledNumsChan {
		fmt.Println(num)
	}
}
