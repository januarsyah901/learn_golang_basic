package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Call() {
	man.Name = "Mr." + man.Name
}

func main() {
	janu := Man{"Januarsyah"}
	janu.Call()
	fmt.Println(janu)
}