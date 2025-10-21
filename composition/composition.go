package main

/*
struct embedding is used to achieve composition
inheritance vs composition
*/

import "fmt"

type Speaker interface {
	Speak() string
}

type Animal struct {
	Legs int
	Name string
}

func NewAnimal(legs int, name string) Speaker {
	return &Animal{
		Legs: legs,
		Name: name,
	}
}

func (an *Animal) Speak() string {
	return "Animal Speaking"
}

func (an *Animal) IsAnimal() bool {
	return true
}

type Dog struct {
	Animal
}

func NewDog() Speaker {
	animal := NewAnimal(4, "dog")
	a, _ := animal.(*Animal)

	return &Dog{
		Animal: *a,
	}
}

func (d *Dog) IsDog() bool {
	if d.Name == "dog" {
		return true
	}
	return false
}

func (an *Dog) Speak() string {
	return "Dog Barking"
}

type Cat struct {
	Animal
}

func NewCat() Speaker {
	animal := NewAnimal(4, "cat")
	a, _ := animal.(*Animal)

	return &Cat{
		Animal: *a,
	}
}

func (c *Cat) IsCat() bool {
	if c.Name == "cat" {
		return true
	}
	return false
}

func TryPolymorphismAndCompositionScenarios() {
	// Check the interface for polymorphism
	// New Dog
	dog := NewDog()
	fmt.Println(dog.Speak())
	realDog := dog.(*Dog)
	fmt.Print(realDog.IsDog())

}
