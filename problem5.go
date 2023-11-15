package main

import "fmt"

func pangkat(base int, pangkat int) int {
	result := 1

	for i := 0; i < int(pangkat); i++ {
		result *= base
	}

	return result
}

func main() {

	fmt.Println(pangkat(2, 3))
}
