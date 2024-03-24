package main

import (
	"fmt"
)

func partition(arr []int, low, high int) ([]int, int) {
	// Set pivot
	pivot := arr[high]
	i := low

	// We need to virtually cut array on two parts. We need a
	// pivot index as divider, so we increment it each iteration,
	// when swap elements (lower than pivot - left, higher - right)
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			// Swap elements
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSortWrapper(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSortWrapper(arr, low, p-1)
		arr = quickSortWrapper(arr, p+1, high)
	}
	return arr
}

// Wrapper for quicksort func
func quicksort(arr []int) []int {
	return quickSortWrapper(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{234, 645, 1, 65, 787, 12, 43, 54, 574, 89, 23, 214, 2}

	arr = quicksort(arr)

	fmt.Println(arr)
}
