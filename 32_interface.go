package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func  (person Person) GetName () string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	janu := Person{"Janu"}
	SayHello(janu)
	kucing := Animal{"Kucing"}
	SayHello(kucing) 
}