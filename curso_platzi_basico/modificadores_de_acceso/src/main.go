package main

import (
	"fmt"
	"my_package/my_package"
)

func main() {
	fmt.Println("MODIFICADORES DE ACCESO")

	var myCar my_package.CarPublic
	myCar.Brand = "Toyota"
	myCar.Model = "Corolla"
	myCar.Year = 2020
	fmt.Println(myCar)

	my_package.PrintMessage()
}

// I had to run   go mod init my_package  to create the module
// Then I had to run go run main.go to run the program
