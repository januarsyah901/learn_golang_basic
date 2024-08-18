package main

import "fmt"

func main() {
	name := map[string]string{
		"firstname": "Januarsyah",
		"middlename": "",
		"lastname":  "akbar",
	}

	fmt.Println(name["firstname"])

	// map function
	// delete
	delete(name, "middlename")
	fmt.Println(name)
	// len
	fmt.Println(len(name))
	

}
