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
	andi := Customers{"Andi", "Ahmad Yani", "Jakarta", "Indonesia", 25}
	// duplikat
	surya := andi
	surya.Name = "Surya"
	surya.City = "Surabaya"
	fmt.Println(surya)
	fmt.Println(andi)

	// pointer
	joko := &andi
	joko.Name = "Joko"
	joko.City = "Bandung"
	fmt.Println(joko)
	fmt.Println(andi)

}
