// This program register the corresponding getter and setter for x and y axis
// x: getX, setX
// y: getX, setY

package main

import "fmt"

type Getter = func() int
type Setter = func(int)

// x live in the scope of getter, setter, x_getter_setter
func axisGetterSetter() (Getter, Setter) {
	var i int

	// Register getter function
	// Use closure to close variable i to both getter adn setter
	// Set them share the same variable
	getter := func() int {
		return i
	}
	setter := func(n int) {
		i = n
	}

	// Return the corresponding get and set function to user
	return getter, setter
}
func main() {
	// register getX  setX to get and set x axis
	getX, setX := axisGetterSetter()

	// register getY, setY to get and set y axis
	getY, setY := axisGetterSetter()

	// Set x axis
	setX(20)
	// Get the x axis value
	fmt.Println("x: ", getX())

	// Set y axis
	setY(5)
	// Get the y axis value
	fmt.Println("y: ", getY())

	// Set x axis again
	setX(100)
	// Get the x axis value
	fmt.Println("x: ", getX())

	// Set y axis again
	setY(0)
	// Get the y axis value
	fmt.Println("y: ", getY())

}
