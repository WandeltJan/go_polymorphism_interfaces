// Golang program to illustrate the
// concept of interfaces
package main

import "fmt"

// interfaces
type car interface {
	speed() int
	name() string
}

type fourWheeler interface{
	engine() string
}

// structs
type SUV struct {
	maxSpeed int
	enginetype string
}

type Combi struct {
	maxSpeed int
}

// functions
func (suv SUV) speed() int {
	return suv.maxSpeed
}

func (suv SUV) name() string {
	return "SUV"
}

func (suv SUV) engine() string{
	return suv.enginetype
}

func (combi Combi) speed() int {
	return combi.maxSpeed
}

func (combi Combi) name() string {
	return "Combi"
}

// main function
func main() {
	//initialize cars
	car1 := Combi{
		maxSpeed: 220,
	}

	car2 := SUV{
		maxSpeed: 130,
	}

	allcars := []car{
		car1,
		car2,
	}

	for _, car := range allcars{
		fmt.Printf("%s has a maximum speed of: %d km / h \n",car.name(), car.speed())
	}

	cars := make(chan car)

	go func() {cars <- car1}()
	go func() {cars <- car2}()

	fmt.Printf("%s was in the channel\n", (<-cars).name() )
	fmt.Printf("%s was in the channel\n", (<-cars).name() )

	fourWheel := make(chan fourWheeler)

	fourWheel <- car2
	//fourWheel <- car1

}