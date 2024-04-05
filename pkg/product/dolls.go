package product

import "fmt"

// Doll is the product interface for a type of toy.
type Doll interface {
	PlayWith()
}

// ClassicDoll is a concrete product that implements Doll.
type ClassicDoll struct{}

// PlayWith allows to play with the ClassicDoll.
func (d *ClassicDoll) PlayWith() {
	fmt.Println("Playing with a classic doll.")
}

// ModernDoll is a concrete product that implements Doll.
type ModernDoll struct{}

// PlayWith allows to play with the ModernDoll.
func (d *ModernDoll) PlayWith() {
	fmt.Println("Playing with a modern doll.")
}
