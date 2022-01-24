// Golang program to illustrate the
// concept of interfaces
package main

import (
	"fmt"
	"math"
)

// interfaces
type twoD_Figure interface {
	Area() float64
}

type threeD_Figure interface {
	Volume() float64
	Surface() float64
}

// 2D structs
type Square struct {
	side float64
}

type Rectangle struct {
	length float64
	width float64
}

type Circle struct {
	radius float64
}

// 3D structs
type Dice struct {
	side Square
}

// functions
func (sqr Square) Area() float64{
	return sqr.side * sqr.side
}

func (rec Rectangle) Area() float64 {
	return rec.width * rec.length
}

func (cir Circle) Area() float64 {
	return cir.radius * cir.radius * math.Pi
}

func (dice Dice) Volume() float64{
	return dice.side.side * dice.side.side * dice.side.side
}

func (dice Dice) Surface() float64{
	return 6* dice.side.Area()
}

// main function
func main() {

	square := Square{

		side: 15.0,
	}

	rectangle := Rectangle{

		length: 10.5,
		width: 12.25,
	}

	circle := Circle{
		radius: 10,
	}

	fmt.Printf("Area of square: \t%.3f unit sq.\n", square.Area())
	fmt.Printf("Area of rectangle: \t%.3f unit sq.\n", rectangle.Area())
	fmt.Printf("Area of circle: \t%.3f unit sq.\n", circle.Area())

	dice := Dice {
		side: square,
	}

	fmt.Printf("Volume of dice \t\t%.3f units cu.\n", dice.Volume())
	fmt.Printf("Length of Cube\t\t%3f units",dice.side.side)
}