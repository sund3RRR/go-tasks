package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const taskNum = 6

func setBit(num int64, pos uint, bit uint) int64 {
	mask := int64(1 << pos)
	if bit == 0 {
		mask = ^mask
		return num & mask
	}
	return num | mask
}

func printHelp() {
	fmt.Println("Usage: ./main <num> <pos> <bit>")
	fmt.Println("\t- num is int64")
	fmt.Println("\t- pos is uint from [0, 63]")
	fmt.Println("\t- bit is uint from [0, 1]")
}

func doTask() {
	args := os.Args[1:]

	if len(args) < 3 {
		log.Println("ERROR: not all args are provided")
		printHelp()
		os.Exit(1)
	}

	num, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		log.Printf("ERROR: couldn't parse num\n %s\n", err)
		printHelp()
		os.Exit(1)
	}

	pos, err := strconv.ParseUint(args[1], 10, 6)
	if err != nil {
		log.Printf("ERROR: couldn't parse pos\n %s\n", err)
		printHelp()
		os.Exit(1)
	}
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
	fmt.Printf("\nStart task %d\n\n", taskNum)

	doTask()

	fmt.Printf("\nEnd of the task %d\n", taskNum)
}
