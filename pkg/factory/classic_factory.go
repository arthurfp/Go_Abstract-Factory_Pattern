package factory

import "abstract-factory-go/pkg/product"

// ClassicToyFactory produces toys with a traditional design.
type ClassicToyFactory struct{}

// NewDoll creates a doll in the classic style.
func (f *ClassicToyFactory) NewDoll() product.Doll {
	return &product.ClassicDoll{}
}

// NewCar creates a car reminiscent of vintage models.
func (f *ClassicToyFactory) NewCar() product.Car {
	return &product.ClassicCar{}
}
