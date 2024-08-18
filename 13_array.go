package main;

import "fmt";

func main() {
	
	var (
		nama[3]string // berarti kapasitasnya 3 index bukan  sampai index ke 3
	) // kalo kosong maka string kosong
	nama[0] = "januarysah"
	nama[1] = "akbar"
	nama[2] = "faiz"

	fmt.Println(nama)

	var (
		angka = [3]int{
			1,2,3, // kalo semisal indexnya kosong maka defaultnya nol
		}
	)
	fmt.Println(angka)
	
	var (
		unkey = [...]int{
			1,2,3,4,5, // harus langsung pas dideklarasikan
		}
	)
	// ganti nilai
	unkey[1] = 3

	// function array
	fmt.Println(len(unkey))
	fmt.Println(unkey[1])


	
}