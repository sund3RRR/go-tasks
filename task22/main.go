/*
	Разработать программу, которая перемножает, делит, складывает,
	вычитает две числовых переменных a,b, значение которых > 2^20.
*/

package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Create two big nums
	a := big.NewFloat(34523453462)
	b := big.NewFloat(61423563456)

	// Create add operation
	add := new(big.Float)
	added := add.Add(a, b)
	fmt.Printf("%f + %f = %f\n", a, b, added)

	// Create subtraction operation
	sub := new(big.Float)
	subtracted := sub.Sub(a, b)
	fmt.Printf("%f - %f = %f\n", a, b, subtracted)

	// Create division operation
	div := new(big.Float)
	divided := div.Quo(a, b)
	fmt.Printf("%f / %f = %f\n", a, b, divided)

	// Create multiplication operation
	mul := new(big.Float)
	multiplicated := mul.Mul(a, b)
	fmt.Printf("%f * %f = %f\n", a, b, multiplicated)
}
