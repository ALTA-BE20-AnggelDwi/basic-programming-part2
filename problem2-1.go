package main

import "fmt"

func main() {
	var bilangan int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bilangan)

	fmt.Printf("Faktor dari %d adalah:\n", bilangan)

	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Printf("%d \n", i)
		}
	}
}
