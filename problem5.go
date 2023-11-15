package main

import "fmt"

func pangkat(base float64, pangkat float64) float64 {
	result := 1.0

	for i := 0; i < int(pangkat); i++ {
		result *= base
	}

	return result
}

func main() {

	fmt.Println(pangkat(2, 3))
}
