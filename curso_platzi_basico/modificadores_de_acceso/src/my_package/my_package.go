package my_package

import "fmt"

// PUBLIC, because the first letter is uppercase
type CarPublic struct {
	Brand string // Ensure this is capitalized to be exported / to be public
	Model string
	Year  int
}

// PRIVATE, because the first letter is lowercase
// type car struct {
// 	brand string
// 	model string
// 	year  int
// }

// function print message - PUBLIC
func PrintMessage() {
	fmt.Println("Hello, I'm a function")
}

// function print message - PRIVATE
// func printMessage() {
// 	fmt.Println("Hello, I'm a function")
// }
