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
	user1 := Customers{"Andi", "Ahmad Yani", "Jakarta", "Indonesia", 25}
	user2 := &user1

	user2.City = "Surabaya"
	// fmt.Println(user2)

	// mungkin ini kaya pindah reference aja, yang awalnya ke A sekarang Ke B
	// user2 = &Customers{"Janu", "Ahmad Yani", "Jakarta", "Indonesia", 23}

	// kalo pengen reference awalnya diikutkan value yang baru maka menggunakan operator *
	*user2 = Customers{"Janu", "Ahmad Yani", "Jakarta", "Indonesia", 23}
	fmt.Println(user1)
	fmt.Println(user2)
}