package main

import (
	"fmt"
	"reflect"
)

func showType(a interface{}) {
	// We use type assertions to determine the type of a variable
	_, ok := a.(string)
	if ok {
		fmt.Println("This is a string")
		return
	}

	_, ok = a.(int)
	if ok {
		fmt.Println("This is an integer")
		return
	}

	_, ok = a.(bool)
	if ok {
		fmt.Println("This is a bool")
		return
	}

	// Use reflect lib to determine that this is a channel because
	// type assertions require to set chan type so we can't use it.
	c := reflect.TypeOf(a)
	if c.Kind() == reflect.Chan {
		fmt.Println("This is a channel")
		return
	}

	fmt.Println("Unknown type!")
}

func main() {
	integer := 343
	showType(integer)

	str := "this is a string"
	showType(str)

	boolka := true
	showType(boolka)

	channelInt := make(chan int)
	showType(channelInt)

	channelInterface := make(chan interface{})
	showType(channelInterface)

	// Unknown type
	float := 34.56
	showType(float)
}
