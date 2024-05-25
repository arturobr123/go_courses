package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) bool {
	array := strings.Split(text, "") // Splitting the string into a slice of its characters
	fmt.Println(array)

	length := len(array)

	for i := 0; i < length/2; i++ {
		if array[i] != array[length-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("RECORRIDO DE SLICES")

	slice := []string{"hola", "mundo", "como", "estas"}

	for index, valor := range slice {
		fmt.Println(index, valor)
	}

	fmt.Println(isPalindromo("oso"))
}
