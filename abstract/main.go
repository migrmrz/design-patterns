package main

import "fmt"

// Animal is the type for our abstract factory.
type Animal interface {
	Says()
	LikesWater() bool
}

// Dog is the concrete factory for dogs.
type Dog struct{}

func (d *Dog) Says() {
	fmt.Println("woof")
}

func (d *Dog) LikesWater() bool {
	return true
}

// Cat is the concrete factory for dogs.
type Cat struct{}

func (d *Cat) Says() {
	fmt.Println("meow")
}

func (d *Cat) LikesWater() bool {
	return true
}

type AnimalFactory interface {
	New() Animal
}

type DogFactory struct{}

func (df *DogFactory) New() Animal {
	return &Dog{}
}

type CatFactory struct{}

func (cf *CatFactory) New() Animal {
	return &Cat{}
}

func main() {
	// create one each of a DogFactory and a CatFactory
	dogFactory := DogFactory{}
	catFactory := CatFactory{}

	// Call new method to create a Dog and a Cat
	dog := dogFactory.New()
	cat := catFactory.New()

	dog.Says()
	cat.Says()

	fmt.Println("A dog likes water:", dog.LikesWater())
	fmt.Println("A cat likes water:", cat.LikesWater())
}
