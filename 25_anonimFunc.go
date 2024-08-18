package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("Maaf, Anda tidak bisa mendaftar.")
	} else {
		fmt.Println("Selamat datang!", name)
	}
}

func main() {
	kataKasar := []string{"anjing", "babi", "monyet"}
	blackList := func(name string) bool {
		for _, kata := range kataKasar {
			if name == kata {
				return true
			}
		}
		return false
	}
	registerUser("anjing", blackList)
	registerUser("januarsyah", blackList)
	registerUser("babi", func(name string) bool {
		return name == "babi"
	})

}
