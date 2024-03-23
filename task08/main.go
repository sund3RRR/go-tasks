package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// The function replace one bit at specified position in num
func setBit(num int64, pos uint, bit uint) int64 {
	// Create bit mask with offset equals bit position
	mask := int64(1 << pos)
	if bit == 0 {
		// Invert mask if bit is 0
		mask = ^mask
		// Bitwise AND operation will set bit at position to zero
		return num & mask
	}

	// Otherwise if bit is 1 just do bitwise OR operation and set
	// bit at position to one
	return num | mask
}

// Help for user if input is incorrect
func printHelp() {
	fmt.Println("Usage: ./main <num> <pos> <bit>")
	fmt.Println("\t- num is int64")
	fmt.Println("\t- pos is uint from [0, 63]")
	fmt.Println("\t- bit is uint from [0, 1]")
}

func doTask() {
	args := os.Args[1:]

	// Check user input is correct
	if len(args) < 3 {
		log.Println("ERROR: not all args are provided")
		printHelp()
		os.Exit(1)
	}

	// Parse num
	num, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		log.Printf("ERROR: couldn't parse num\n %s\n", err)
		printHelp()
		os.Exit(1)
	}

	// Parse bit position
	pos, err := strconv.ParseUint(args[1], 10, 6)
	if err != nil {
		log.Printf("ERROR: couldn't parse pos\n %s\n", err)
		printHelp()
		os.Exit(1)
	}

	// Parse bit
	bit, err := strconv.ParseUint(args[2], 10, 1)
	if err != nil {
		log.Printf("ERROR: couldn't parse bit\n %s\n", err)
		printHelp()
		os.Exit(1)
	}

	result := setBit(num, uint(pos), uint(bit))
	fmt.Printf("num = %d\n", result)
}

func main() {
	doTask()
}
