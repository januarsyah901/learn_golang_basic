package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}
func runApp(error bool) {
	defer logging()
	if error {
		panic("Error mas!")
	}
	fmt.Println("Run application")
	
}
func main() {
	runApp(true)
}