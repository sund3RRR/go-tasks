/*
	Реализовать бинарный поиск встроенными методами языка.
*/

package main

import "fmt"

// Binary search is a simple but really powerful algorithm for searching
// items in sorted arrays. The main idea is to cut array in a half each iteration
// of algorithm, so it means that the time complexity is O(log(n)).
func binarySearch(arr []int, target int) int {
	// Create to pointers
	left, right := 0, len(arr)-1

	for left <= right {
		// Get mid index between two pointers
		mid := (left + right) / 2

		if arr[mid] < target {
			// Cut left half
			left = mid + 1
		} else if arr[mid] > target {
			// Cut right half
			right = mid - 1
		} else {
			return mid
		}
	}

	// We didn't find element, so return -1
	return -1
}

func main() {
	arr := []int{1, 2, 4, 6, 8, 34, 576, 874}

	idx := binarySearch(arr, 34)
	if idx == -1 {
		fmt.Println("There is no such element in array :(")
		return
	}
	fmt.Println("Element index is", idx)
}
