package main

import "fmt"

func main() {
	// for 
	count := 0
	for count <= 10 {
		fmt.Println("Perulangan ke-", count)
		count++
	}
	// for statment
	for counter := 0; counter <=12; counter++ {
		fmt.Println("Perulangan ke-", counter)
	}
	// for range
	name := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}

	for index, name := range name{
		fmt.Printf("index : %v, name : %v \n", index, name)
	}
	// kalo indexnya ga butuh tinggal ganti " _ "
}

