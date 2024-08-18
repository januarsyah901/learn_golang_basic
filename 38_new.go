package main

import "fmt"

type Person struct {
	name, address string
	age          int
}

func main() {
	person1 := new(Person)
	person2 := person1
	person2.name = "Januarsyah"
	fmt.Println(person1.name)
	fmt.Println(person2.name)
}