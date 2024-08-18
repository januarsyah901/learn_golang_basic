package main

import "fmt"

type Customers struct {
	Name, Strict, City, Country string
	Age                         int
	
}

// methode say hello
func (customer Customers) sayHello(name string) {
	fmt.Println("Hello",name,", my name is", customer.Name)
}

func main() {
	// cara 1
	var janu Customers 
	janu.Name = "Janu"
	janu.Strict = "Ahmad Yani"
	janu.City = "Jakarta"
	janu.Country = "Indonesia"
	janu.Age = 25
	janu.sayHello("Andi")

	// cara 2
	surya := Customers {
		Name: "Surya",
		Strict: "Ahmad Yani",
		City: "Jakarta",
		Country: "Indonesia",
		Age: 25,
	}

	// cara 3
	andi := Customers{"Andi", "Ahmad Yani", "Jakarta", "Indonesia", 25}



	fmt.Println(janu)
	fmt.Println(surya)
	fmt.Println(andi)
}
