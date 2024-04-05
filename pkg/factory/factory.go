package factory

import "abstract-factory-go/pkg/product"

// ToyFactory is the abstract factory interface that defines methods for creating different types of toys.
type ToyFactory interface {
	NewDoll() product.Doll
	NewCar() product.Car
}
