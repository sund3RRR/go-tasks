/*
	Разработать программу нахождения расстояния между двумя точками,
	которые представлены в виде структуры Point с инкапсулированными
	параметрами x,y и конструктором.
*/

package main

import "fmt"

// Just find distance between two points
func Distance(a, b *Point) float64 {
	c := a.SubtractVector(b)

	return c.Magnitude()
}

func main() {
	a := NewPoint(2, 7)
	b := NewPoint(45, 12)

	fmt.Println("Distance between two points:", Distance(a, b))
}
