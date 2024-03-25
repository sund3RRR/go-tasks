/*
	Дана структура Human (с произвольным набором полей и методов).
	Реализовать встраивание методов в структуре Action от родительской
	структуры Human (аналог наследования).
*/

package main

import "fmt"

type Human struct {
	Name        string
	Surname     string
	Age         int
	Nationality string
}

type Action struct {
	Human
}

func (h *Human) SayHello() {
	fmt.Printf("Hi, my name is %s %s, I'm %d years old and I'm %s\n",
		h.Name, h.Surname, h.Age, h.Nationality)
}

func (a *Action) Run(steps int) {
	fmt.Printf("Running %d steps\n", steps)
}

func main() {
	action := Action{
		Human{
			Name:        "John",
			Surname:     "Silver",
			Age:         49,
			Nationality: "British",
		},
	}

	action.SayHello()
	action.Run(20)
}
