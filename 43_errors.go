package main

import (
	"errors"
	"fmt"
)

func filterWord(word string) (string, error) {
	if word == "anjing" {
		return "", errors.New("kata ini tidak diizinkan")
	}
	return word, nil
}

func main() {
	result, err := filterWord("anjing")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}