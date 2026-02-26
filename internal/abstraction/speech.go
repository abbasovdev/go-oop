package abstraction

import "fmt"

// Speaker is an interface that defines a single behavior: Speak().
// Abstraction: We care about what something can do (speak), not how it's implemented.
// In Go, interfaces are implicit: no "implements" keyword; any type with these methods satisfies the interface.
type Speaker interface {
	Speak() string
}

// Dog is a concrete type that can speak. It implements Speaker by having a Speak() method.
type Dog struct {
	Name string
}

// Speak returns what a dog says. This satisfies the Speaker interface.
func (d Dog) Speak() string {
	return fmt.Sprintf("%s says: Woof!", d.Name)
}

// Robot is another concrete type that can "speak". Same interface, different implementation.
// Abstraction: callers use Speaker and don't need to know if it's a Dog or Robot.
type Robot struct {
	Model string
}

// Speak returns what a robot says. Robot also satisfies Speaker.
func (r Robot) Speak() string {
	return fmt.Sprintf("Robot %s says: Beep boop.", r.Model)
}
