package main

import (
	"fmt"
)

func main() {
	temp := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Create map where key - int and value - slice of float32 temps
	hashMap := make(map[int][]float32)

	for _, tmp := range temp {
		// Get key for map
		key := int(float64(tmp)/10) * 10

		// Append temp to group
		hashMap[key] = append(hashMap[key], tmp)
	}

	fmt.Println(hashMap)
}
