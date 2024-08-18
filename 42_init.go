package main

import (
	"learn-golang-basic/databases"
	_"learn-golang-basic/internal"
	"fmt"
)

func main() {
	fmt.Println(databases.GetConnection())
}