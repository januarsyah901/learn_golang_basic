package main

import "fmt"

func goodBye(name string) string {
	return "GoodBye " + name
}

func main() {
	contoh := goodBye
	fmt.Println(contoh("januarsyah"))
}