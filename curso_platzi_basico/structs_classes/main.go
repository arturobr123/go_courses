package main

import "fmt"

type car struct {
	brand string
	model string
	year  int
}

func main() {
	fmt.Println("STRUCTS AND CLASSES")

	mycar := car{brand: "Toyota", model: "Corolla", year: 2020}
	fmt.Println(mycar)

	var otherCar car
	otherCar.brand = "Ford"
	otherCar.model = "Fusion"
	// year will be 0, because it is the 0 value of the int type
	fmt.Println(otherCar)
}
