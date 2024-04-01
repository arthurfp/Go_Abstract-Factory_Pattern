package factory

// ToyFactory is the abstract factory interface.
type ToyFactory interface {
	NewDoll() Doll
	NewCar() Car
}
