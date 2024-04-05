package factory

import "abstract-factory-go/pkg/product"

// ModernToyFactory is a concrete factory that produces modern toys.
type ModernToyFactory struct{}

// NewDoll returns a new instance of a ModernDoll.
func (f *ModernToyFactory) NewDoll() product.Doll {
	return &product.ModernDoll{}
}

// NewCar returns a new instance of a ModernCar.
func (f *ModernToyFactory) NewCar() product.Car {
	return &product.ModernCar{}
}
