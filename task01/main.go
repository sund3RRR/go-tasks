package main

import "fmt"

const taskNum = 1

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
	fmt.Printf("\nStart task %d\n\n", taskNum)

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

	fmt.Printf("\nEnd of the task %d\n", taskNum)
}
