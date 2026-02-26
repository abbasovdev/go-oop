package inheritance

import "fmt"

// Engine is a reusable component. In traditional OOP this might be a "base class."
// Inheritance: We use composition (embedding) instead of class inheritance.
type Engine struct {
	Horsepower int
	Running    bool
}

// Start turns the engine on. Embedded structs' methods are promoted to the outer type.
func (e *Engine) Start() {
	e.Running = true
	fmt.Printf("Engine started (%d HP).\n", e.Horsepower)
}

// Stop turns the engine off.
func (e *Engine) Stop() {
	e.Running = false
	fmt.Println("Engine stopped.")
}

// Car embeds Engine. This is struct embedding (composition).
// Go way: Car "has an" Engine; we don't say Car "is an" Engine.
// The embedded Engine's fields and methods are promoted, so Car gets Start(), Stop(), etc.
type Car struct {
	Engine // embedded anonymously, promotes Engine's fields and methods
	Model  string
	Wheels int
}

// NewCar constructs a Car with an embedded Engine. Demonstrates composition in use.
func NewCar(model string, horsepower int) Car {
	return Car{
		Engine: Engine{Horsepower: horsepower, Running: false},
		Model:  model,
		Wheels: 4,
	}
}

// Drive uses the embedded Engine. We can call c.Start() because Engine is embedded.
func (c *Car) Drive() {
	if !c.Running {
		c.Start()
	}
	fmt.Printf("%s is driving on %d wheels.\n", c.Model, c.Wheels)
}
