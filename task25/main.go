package main

import (
	"fmt"
	"time"
)

// The function will sleep for the specified time in seconds
func sleep(seconds int) {
	// time.After() creates a pipe that sends a value after a specified time,
	// and the <- operator blocks execution until the value is received from the pipe
	<-time.After(time.Duration(seconds) * time.Second)
}

func main() {
	const SLEEP_TIME_SEC = 5

	fmt.Printf("Starting to sleep for %d seconds\n%s\n", SLEEP_TIME_SEC, time.Now())
	sleep(SLEEP_TIME_SEC)
	fmt.Println("Awake", time.Now())
}
