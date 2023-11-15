package main

import "fmt"

// Fungsi untuk memeriksa apakah suatu bilangan prima penuh atau tidak
func isFullPrime(num int) bool {
	if num < 2 {
		return false
	}

	// Fungsi untuk memeriksa apakah suatu bilangan prima
	isPrime := func(n int) bool {
		if n < 2 {
			return false
		}
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	// Memeriksa apakah bilangan itu sendiri adalah prima
	if !isPrime(num) {
		return false
	}

	// Memeriksa apakah semua digit juga prima
	digits := num
	for digits > 0 {
		digit := digits % 10
		if !isPrime(digit) {
			return false
		}
		digits /= 10
	}

	return true
}

func main() {
	// Contoh penggunaan
	fmt.Println(isFullPrime(23))
	fmt.Println(isFullPrime(2231))
}
