/*
	Разработать программу, которая переворачивает слова в строке.
	Пример: «snow dog sun — sun dog snow».
*/

package main

import (
	"fmt"
	"log"
	"strings"
)

func invertWords(str string) string {
	// Create string builder
	builder := strings.Builder{}

	// Split string to wordss
	splitted := strings.Split(str, " ")

	for i := len(splitted) - 1; i >= 0; i-- {
		// Write string to builder and add space separator
		_, err := builder.WriteString(splitted[i] + " ")
		if err != nil {
			log.Fatal(err)
		}
	}

	// Trim the last space and build string
	return strings.TrimSuffix(builder.String(), " ")
}

func main() {
	str := "snow dog sun"
	fmt.Println(str)

	str = invertWords(str)
	fmt.Println(str)
}
