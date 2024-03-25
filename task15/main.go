/*
	К каким негативным последствиям может привести данный фрагмент
	кода, и как это исправить? Приведите корректный пример реализации.


	var justString string
	func someFunc() {
		v := createHugeString(1 << 10)
		justString = v[:100]
	}

	func main() {
		someFunc()
	}
*/

package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
)

/*
var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}

This code snippet has a potential memory leak issue.
The someFunc() function creates a large string v,
and then extracts a substring from it into the justString variable.
However, the substring justString retains a reference to the original large string v,
which may cause the large string to not be collected by the garbage collector even if it is no longer needed.
This may result in a memory leak.

To fix this, instead of saving the substring, we should create a new string containing only the desired data.
*/

var justString string

func createHugeString(length int) string {
	// Create bytes buffer to build string
	var buffer bytes.Buffer
	for length > 0 {

		// Write rune to buffer
		_, err := buffer.WriteRune('a')
		if err != nil {
			log.Fatal("ERROR: ", err)
		}
		length--
	}

	// Build string
	return buffer.String()
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()

	fmt.Println(justString)
}
