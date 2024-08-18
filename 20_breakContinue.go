package main

import "fmt"

func main() {
    // Contoh penggunaan break dalam loop
    for i := 1; i <= 10; i++ {
        if i == 5 {
            break // keluar dari loop ketika i sama dengan 5
        }
        fmt.Println("Iterasi ke-", i)
    }

    // Contoh penggunaan continue dalam loop
    for j := 1; j <= 10; j++ {
        if j%2 == 0 {
            continue // skip iterasi saat j adalah bilangan genap
        }
        fmt.Println("Iterasi ke-", j)
    }
   
}
