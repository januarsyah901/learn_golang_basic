package main

import "fmt"

func main() {
	counter := 0
	increment := func() {
		fmt.Println("Increment")
		counter++
	}
	decrement := func() {
		fmt.Println("Decrement")
		counter--
	}
	increment()
	increment()
	fmt.Println(counter)
	decrement()
	fmt.Println(counter)
}