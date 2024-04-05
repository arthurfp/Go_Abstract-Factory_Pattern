package factory

import "abstract-factory-go/pkg/product"

// ClassicToyFactory is a concrete factory that produces classic toys.
type ClassicToyFactory struct{}

// NewDoll returns a new instance of a ClassicDoll.
func (f *ClassicToyFactory) NewDoll() product.Doll {
	return &product.ClassicDoll{}
}

// NewCar returns a new instance of a ClassicCar.
func (f *ClassicToyFactory) NewCar() product.Car {
	return &product.ClassicCar{}
}
