package main

import (
	"abstract-factory-go/pkg/factory"
	"fmt"
)

func main() {
	fmt.Println("Abstract Factory Pattern in Go")

	// Create a ClassicToyFactory
	classicFactory := factory.ClassicToyFactory{}

	// Create and use a Doll and Car from ClassicToyFactory
	doll := classicFactory.NewDoll()
	doll.PlayWith()
	car := classicFactory.NewCar()
	car.Drive()

	// Create a ModernToyFactory
	modernFactory := factory.ModernToyFactory{}

	// Create and use a Doll and Car from ModernToyFactory
	modernDoll := modernFactory.NewDoll()
	modernDoll.PlayWith()
	modernCar := modernFactory.NewCar()
	modernCar.Drive()
}
