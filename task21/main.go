package main

import "fmt"

// Internal base interface implementation
type InputDevice interface {
	Click()
}

//
// Third party object
//

type Touchpad struct {
}

func (t *Touchpad) Tap() {
	fmt.Println("**tap**")
}

//
// Third party object
//

//
// Adapter that converts Touchpad methods to InputDevice methods
//

type InputAdapter struct {
	touchpad *Touchpad
}

func (a *InputAdapter) Click() {
	fmt.Println("**click**")
	a.touchpad.Tap()
}

//
// Adapter that converts Touchpad methods to InputDevice methods
//

func main() {
	a := InputAdapter{}
	a.Click()
}
