package main

import "fmt"

func Area(length int, breadth int) int {
	return length * breadth
}

func Area(side int) int {
	return side * side
}

func main() {
	fmt.Println(Area(5))
}

