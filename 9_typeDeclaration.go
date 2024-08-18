package main;

import "fmt";

func main() {
	
	type noKTP string

	var KTPku noKTP = "11111111"

	var contoh string = "2222222"

	var contohKTP noKTP = noKTP(contoh)

	fmt.Println(KTPku)
	fmt.Println(contohKTP)

}