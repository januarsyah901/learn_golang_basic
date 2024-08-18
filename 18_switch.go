 // Start Generation Here
package main

import "fmt"

func main() {
    var name string
	name = "januarsyahakbar"

    switch name {
    case "januarsyah":
        fmt.Println("Halo Januarsyah!")
    case "surya":
        fmt.Println("Halo Surya!")
    default:
        fmt.Println("Halo, " + name + "!")
	}
	// switch with short statement
	switch name_len := len(name); name_len <= 15  {
	case true:
		fmt.Println("pas")
	case false:
		fmt.Println("nama kepanjangan bre")
	}

	// next switch
	switch name_len := len(name){
	case name_len >= 15:
		fmt.Println("namanya pas")
	case name_len >= 20:
		fmt.Println("nama kepanjangan bre")
	default:
		fmt.Println("wedeh ga normal")
	}
}
