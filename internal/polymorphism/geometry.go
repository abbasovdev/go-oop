package polymorphism

import (
	"fmt"
	"math"
)

// Shape is an interface that defines behavior for geometric shapes.
// Polymorphism: Different types (Circle, Square) can be used interchangeably
// wherever a Shape is needed. Go achieves this via implicit interface satisfaction.
type Shape interface {
	Area() float64
	Perimeter() float64
	String() string
}

// Circle implements Shape. No "implements" keyword, having the methods is enough.
type Circle struct {
	Radius float64
}

// Area returns π*r². This method makes Circle satisfy Shape.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter returns 2πr.
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// String returns a description. Satisfies Shape and fmt.Stringer.
func (c Circle) String() string {
	return fmt.Sprintf("Circle(radius=%.2f)", c.Radius)
}

// Square implements Shape. Same interface, different implementation.
type Square struct {
	Side float64
}

// Area returns side².
func (s Square) Area() float64 {
	return s.Side * s.Side
}

// Perimeter returns 4*side.
func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

// String returns a description.
func (s Square) String() string {
	return fmt.Sprintf("Square(side=%.2f)", s.Side)
}

// PrintShapeInfo demonstrates polymorphism: it accepts any Shape (Circle or Square)
// and calls Area/Perimeter without knowing the concrete type.
func PrintShapeInfo(shapes ...Shape) {
	for _, s := range shapes {
		fmt.Printf("%s — Area: %.2f, Perimeter: %.2f\n", s.String(), s.Area(), s.Perimeter())
	}
}
