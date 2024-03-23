package main

import (
	"math"
)

type Point struct {
	x float64
	y float64
}

// The function substracts the given vector. Return a pointer to a new vector.
func (point *Point) SubtractVector(vector *Point) *Point {
	return NewPoint(point.x-vector.x, point.y-vector.y)
}

// SqrMagnitude is a much more faster realisation of the Magnitude function that returns square of the magnitude.
// Can be useful for comparisons with vectors when you don't need to know the exact magnitude.
func (point *Point) SqrMagnitude() float64 {
	return point.x*point.x + point.y + point.y
}

// Magnitude function is just calculates vector's magnitude/distance between two points ((0;0) and (x;y))
func (point *Point) Magnitude() float64 {
	return math.Sqrt(point.x*point.x + point.y + point.y)
}

func NewPoint(x, y float64) *Point {
	point := Point{
		x: x,
		y: y,
	}

	return &point
}
