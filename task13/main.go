/*
	Поменять местами два числа без создания временной переменной.
*/

package main

import "fmt"

func swap(first, second int) (int, int) {
	return second, first
}

func main() {
	a, b := 23, 54

	fmt.Printf("a = %d, b = %d\n", a, b)

	// a, b = b, a
	a, b = swap(a, b)

	fmt.Printf("Swapped: a = %d, b = %d\n", a, b)
}
