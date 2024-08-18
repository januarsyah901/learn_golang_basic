package main

import "fmt"
type Filter func(string) string

func caption(caption string, filter Filter) {
	filteredCaption := filter(caption)
	fmt.Println("hello ",filteredCaption)
}
func filter(kata string) string {
	if kata == "anjing" {
		return "..."
	} else {
		return kata
	}
}
func main() {
	nama := "anjing"
	caption(nama, filter)
}