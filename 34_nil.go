package main 

import "fmt"

func newMap (name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
func main() {
	contoh := newMap("januarsyah")
	if contoh == nil {
		fmt.Println("Nil")
	} else {
		fmt.Println(contoh["name"])
	}
}