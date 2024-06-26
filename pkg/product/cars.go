package product

import "fmt"

// Car is the product interface for another type of toy.
type Car interface {
	Drive()
}

// ClassicCar is a concrete product that implements Car.
type ClassicCar struct{}

// Drive allows you to pretend to drive the ClassicCar.
func (c *ClassicCar) Drive() {
	fmt.Println("Driving a classic car.")
}

// ModernCar is a concrete product that implements Car.
type ModernCar struct{}

// Drive allows you to pretend to drive the ModernCar.
func (c *ModernCar) Drive() {
	fmt.Println("Driving a modern car.")
}
