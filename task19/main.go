package main

import (
	"fmt"
	"strings"
)

func invertString(str string) string {
	// Convert string to the slice of runes.
	// It is required because Unicode symbols can be presented in
	// multiple count of bytes so we can't iterate over string bytes
	// directly and reverse it because we can break string.
	s := []rune(str)

	// Create string builder
	builder := strings.Builder{}

	for i := len(s) - 1; i >= 0; i-- {
		// Write rune to builder
		builder.WriteRune(s[i])
	}

	// Build string
	return builder.String()
}

func main() {
	str := "главрыба"
	fmt.Println(str)

	str = invertString(str)
	fmt.Println(str)
}
