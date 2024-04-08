package factory

import (
	"abstract-factory-go/pkg/product"
	"reflect"
	"testing"
)

func TestClassicToyFactory(t *testing.T) {
	factory := ClassicToyFactory{}

	doll := factory.NewDoll()
	if reflect.TypeOf(doll) != reflect.TypeOf(&product.ClassicDoll{}) {
		t.Errorf("NewDoll() should return a ClassicDoll, got %T", doll)
	}

	car := factory.NewCar()
	if reflect.TypeOf(car) != reflect.TypeOf(&product.ClassicCar{}) {
		t.Errorf("NewCar() should return a ClassicCar, got %T", car)
	}
}

func TestModernToyFactory(t *testing.T) {
	factory := ModernToyFactory{}

	doll := factory.NewDoll()
	if reflect.TypeOf(doll) != reflect.TypeOf(&product.ModernDoll{}) {
		t.Errorf("NewDoll() should return a ModernDoll, got %T", doll)
	}

	car := factory.NewCar()
	if reflect.TypeOf(car) != reflect.TypeOf(&product.ModernCar{}) {
		t.Errorf("NewCar() should return a ModernCar, got %T", car)
	}
}
