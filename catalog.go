package main

import (
	"fmt"
)

//interface
type order interface {
	shipping() float64
	tax() float64
}

type discount interface {
	offer()
	message()
}

// structs
type clothes struct {
	name string
	price float64
	qty float64
}

type download struct{
	name string
	price float64
	qty float64
}

//struct functions
//cloths
func (c clothes) tax() float64{
	return c.price * c.qty * 0.05
}

func (c clothes) shipping() float64{
	return c.qty * 5
}
//downloads
func (d download) tax() float64{
	return d.price * d.qty * 0.07
}

func (d download) shipping() float64{
	return 0
}

func (d download) offer() float64{
	return d.price * 0.5
}

func (d download) message() string{
	return "limited time just until christmas"
}

func main(){
	cl1 := clothes{
		name: "T-shirt",
		price: 10.0,
		qty: 2.0,
	}

	d1 := download{
		name: "Hands-On Software Engineering with Golang: Move beyond basic programming to design and build " +
			"reliable software with clean code Taschenbuch – 24. Januar 2020",
		price: 40.64,
		qty: 42.0,
	}

	allorders := []order{
		cl1,
		d1,
	}

	for _, order := range allorders{
		fmt.Printf("Shipping Charge: %3f\n", order.shipping())
		fmt.Printf("Tax: %3f\n", order.tax())
	}


}