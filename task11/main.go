/*
	Реализовать пересечение двух неупорядоченных множеств.
*/

package main

import "fmt"

// This function finds an intersaction of two integer arrays.
// The algorithm complexity is about O(max(n,m)), but the memory consumption is O(min(n,m)).
func intersection(first, second []int) []int {
	// Swap slices if the first slice in bigger than the second
	if len(first) > len(second) {
		first, second = second, first
	}

	// Create hash map for smaller slice
	hashMap := make(map[int]bool, len(first))

	// Fill hash map with nums from smaller slice
	for _, num := range first {
		hashMap[num] = false
	}

	// Create result slice with intersection
	result := make([]int, 0, 10)

	for _, num := range second {
		// Check that there is an intersection
		if _, exists := hashMap[num]; exists {
			// Check if num is already in result slice
			if hashMap[num] {
				continue
			}
			// Append intersection num into result slice
			result = append(result, num)
			// Mark that we already append this num to result
			hashMap[num] = true
		}
	}

	return result
}

func main() {
	first := []int{35, 14, 64, 867, 42, 1, 42, 65, 33}
	second := []int{54, 76, 14, 35, 2, 42, 42}

	result := intersection(first, second)

	fmt.Println(result)
}
