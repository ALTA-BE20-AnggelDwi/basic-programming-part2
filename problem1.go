package main

import "fmt"

func main() {
	var studentScore int = 40

	if studentScore == 80 && studentScore <= 100 {
		fmt.Println("Nilai mashasiswa: A")
	} else if studentScore == 65 && studentScore < 80 {
		fmt.Println("Nilai mashasiswa: B+")
	} else if studentScore == 50 && studentScore < 65 {
		fmt.Println("Nilai mashasiswa: B")
	} else if studentScore == 35 && studentScore < 50 {
		fmt.Println("Nilai mashasiswa: C")
	} else {
		fmt.Println("Nilai mashasiswa: D")
	}

}
