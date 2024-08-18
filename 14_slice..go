package main

import "fmt"

func main() {
	var huruf = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	// var slice = huruf[1:3]
	// fmt.Println(slice)
	// var slice2 = huruf[3:]
	// fmt.Println(slice2)
	// var slice3 = huruf[:4]
	// fmt.Println(slice3)
	// var slice4 = huruf[:]
	// fmt.Println(slice4)

	// pembagian huruf
	konsonan := huruf[4:]
	fmt.Println(konsonan)

	// mengganti huruf
	konsonan[1] = "Z"
	fmt.Println(konsonan)

	// menambah huruf
	huruf_baru := append(konsonan, "K")
	fmt.Println("huruf_baru: ", huruf_baru)
	fmt.Println("huruf: ", huruf)
	fmt.Println("konsonan: ", konsonan)
}
 
 
 