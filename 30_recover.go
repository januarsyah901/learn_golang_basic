package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}
func errorCheck() {
	massage := recover()
	fmt.Println("terjadi error :", massage)
}
func runApp(error bool) {
	defer logging()
	defer errorCheck()
	if error {
		panic("Error mas!")
	}
	fmt.Println("Run application")
	
}
func main() {
	runApp(true)
	fmt.Println("selesai")
}