package main

import (
	"learn-golang-basic/helper"
	"fmt"
)

func main() {
	name := "januarysah"
	name = helper.SayHello(name)
	fmt.Println(name)
	fmt.Println(helper.Info())
	fmt.Println(helper.Application)
}