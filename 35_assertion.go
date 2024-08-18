package main

import "fmt"

func random () any {
	return true
}
func main() {
	contoh := random()
	// yang bener
	// contohString := contoh.(string)
	// fmt.Println(contohString)

	// yang  salah
	// contohSalah := contoh.(int) // salah]\

	switch value := contoh.(type) {
	case string:
		fmt.Println("String ", value)
	case int:
		fmt.Println("Int ", value)
	default:
		fmt.Println("buh ra ruh ")
	}
	
}