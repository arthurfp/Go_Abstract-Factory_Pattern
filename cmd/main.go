package main

import (
	"abstract-factory-go/pkg/factory"
	"fmt"
)

func main() {
	fmt.Println("Abstract Factory Pattern in Go")

	// Create a ClassicToyFactory
	classicFactory := factory.ClassicToyFactory{}

	// Create a Doll using the ClassicToyFactory
	doll := classicFactory.NewDoll()
	doll.PlayWith()

	// Create a Car using the ClassicToyFactory
	car := classicFactory.NewCar()
	car.Drive()
}
