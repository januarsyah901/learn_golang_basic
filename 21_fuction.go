package main

import "fmt"

func sayHello() {
    fmt.Println("Hello!")
}

func sayHelloTo(name, asal string) {
    fmt.Println("Hello,", name, "kamu dari ", asal, "ya?")
}

func sum(a, b int) int {
    return a + b
}

func multiValues()(string,string,string) {
    return "januarsyah","surya","andi"
}

func main() {
    sayHello()
    sayHelloTo("Januarsyah", "Kalbar,")
	sum := sum(1, 2)
	fmt.Println(sum)
    orang1, orang2, _ := multiValues()
    fmt.Println(orang1, orang2)
}

