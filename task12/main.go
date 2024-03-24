package main

import "fmt"

func contains(arr []string, target string) bool {
	for _, str := range arr {
		// if string is presented in set, return true
		if str == target {
			return true
		}
	}
	return false
}

func makeSet(arr []string) []string {
	set := make([]string, 0)

	for _, str := range arr {
		// If string is not present in set, add it
		if !contains(set, str) {
			set = append(set, str)
		}
	}

	return set
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	set := makeSet(arr)

	fmt.Println(set)
}
