package main

import (
	"fmt"

	"go-oop/internal/abstraction"
	"go-oop/internal/encapsulation"
	"go-oop/internal/inheritance"
	"go-oop/internal/polymorphism"
)

func main() {
	fmt.Println("=== OOP Concepts in Go: Demonstration ===")
	fmt.Println()

	demoEncapsulation()
	fmt.Println()

	demoAbstraction()
	fmt.Println()

	demoInheritance()
	fmt.Println()

	demoPolymorphism()
}

func demoEncapsulation() {
	fmt.Println("--- 1. Encapsulation (internal/encapsulation/bank.go) ---")
	acc := encapsulation.NewAccount(100)
	fmt.Printf("Initial balance: %.2f\n", acc.Balance())
	acc.Deposit(50)
	fmt.Printf("After deposit: %.2f\n", acc.Balance())
	ok := acc.Withdraw(30)
	fmt.Printf("Withdraw 30: ok=%v, balance=%.2f\n", ok, acc.Balance())
	ok = acc.Withdraw(200)
	fmt.Printf("Withdraw 200 (insufficient): ok=%v, balance=%.2f\n", ok, acc.Balance())
}

func demoAbstraction() {
	fmt.Println("--- 2. Abstraction (internal/abstraction/speech.go) ---")
	speakers := []abstraction.Speaker{
		abstraction.Dog{Name: "Rex"},
		abstraction.Robot{Model: "X-100"},
	}
	for _, s := range speakers {
		fmt.Println(s.Speak())
	}
}

func demoInheritance() {
	fmt.Println("--- 3. Inheritance / Composition (internal/inheritance/vehicle.go) ---")
	car := inheritance.NewCar("Sedan", 150)
	car.Drive()
	car.Stop()
}

func demoPolymorphism() {
	fmt.Println("--- 4. Polymorphism (internal/polymorphism/geometry.go) ---")
	shapes := []polymorphism.Shape{
		polymorphism.Circle{Radius: 5},
		polymorphism.Square{Side: 4},
	}
	polymorphism.PrintShapeInfo(shapes...)
}
