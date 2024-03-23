package main

import (
	"fmt"
	"strings"
)

func isUniqueStr(str string) bool {
	// Create hash map for letters
	hashMap := make(map[string]bool, len(str))

	for _, char := range str {
		// Make char lowercase
		key := strings.ToLower(string(char))

		// Check if char is already present in hash map
		if _, exists := hashMap[key]; exists {
			// Then exit and return false, because there is at least one
			// repeating character
			return false
		}

		// Add char to hash map
		hashMap[key] = false
	}

	// No one repeating char, so return true
	return true
}

func main() {
	data := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, str := range data {
		if isUniqueStr(str) {
			fmt.Println("String is unique")
		} else {
			fmt.Println("String has repeating letters")
		}
	}
}
