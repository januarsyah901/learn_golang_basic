package main

import "fmt"

func main() {
	 // if
	 name := "adi asu"
	 if name == "januarsyah" {
		fmt.Println("name is januarsyah")
	 } else if  name == "surya" {
		fmt.Println("hello  surya")
	 } else {
		fmt.Println("who are you?")
	 }
	 // make a if short
	 if name2 := "januarsyah"; name2 == "januarsyah" {
		fmt.Println("name is januarsyah")
	 }
}

