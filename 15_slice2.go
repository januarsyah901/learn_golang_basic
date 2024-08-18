package main

import "fmt"

func main() {
	// make new slice
	var newSlince = make([]string, 2,5)
	newSlince[0] = "A"
	newSlince[1] = "B"
	fmt.Println( len(newSlince))
	newSlince = append(newSlince, "C")
	newSlince = append(newSlince, "D")
	newSlince = append(newSlince, "E")
	fmt.Println(newSlince)
	fmt.Println(len(newSlince))
	fmt.Println(cap(newSlince))
	// copy slice
	fromSlice := newSlince[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	fmt.Println("cap toSlice: ",cap(toSlice))
	copy(toSlice, fromSlice)
	fmt.Println(toSlice)

	// defference slice with array
	array := [...]int{1,2,3,4,5}
	slice := []int{1,2,3,4,5}
	fmt.Println(array)
	fmt.Println(slice)
}
 
 
 