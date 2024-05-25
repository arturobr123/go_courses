package main

import "fmt"

func main() {
	fmt.Println("ARRAYS")

	// array
	var array [5]int // los array son estaticos, no se pueden agregar mas elementos despues de su creacion
	array[0] = 1
	array[1] = 2

	fmt.Println(array)

	// slice
	slice := []int{1, 2, 3, 4, 5} // los slices son dinamicos, se pueden agregar mas elementos despues de su creacion
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice[1:3])

	// append
	slice = append(slice, 6)
	fmt.Println(slice)

	// append new slice
	newSlice := []int{7, 8, 9}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

}
