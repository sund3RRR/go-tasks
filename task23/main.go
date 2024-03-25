/*
	Удалить i-ый элемент из слайса.
*/

package main

import (
	"errors"
	"fmt"
	"log"
)

func removeAt(arr *[]int, idx int) (*[]int, error) {
	// Check if idx is out of range
	if idx < 0 || idx >= len(*arr) {
		return nil, errors.New("slice index out of range")
	}

	// Cut slice on idx+1 index and join using append and slice unpacking
	s := append((*arr)[:idx], (*arr)[idx+1:]...)

	// Return pointer on slice and nil error
	return &s, nil
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}

	s, err := removeAt(&nums, 5)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(*s)
}
