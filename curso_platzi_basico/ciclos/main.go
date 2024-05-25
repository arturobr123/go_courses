package main

import "fmt"

func main() {
	fmt.Println("CICLOS")

	// for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	// while
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println()

	// for forever
	counterForEver := 0
	for {
		fmt.Println("for forever")
		fmt.Println(counterForEver)
		counterForEver++
		if counterForEver == 10 {
			break
		}
	}
}
