package main

import "fmt"

type Number struct {
	Value int
}

func defaultValue(number *Number) {
	number.Value = 0
}

func main() {
	number := Number{}
	number.Value = 10
	defaultValue(&number)
	fmt.Println(number.Value)
}