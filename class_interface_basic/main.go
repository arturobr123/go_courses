// main.go

package main

import "fmt"

func main() {
	c := NewCircle(5)
	printShapeDetails(c)

	r := NewRectangle(3, 4)
	printShapeDetails(r)

	s := NewSquare(5)
	printShapeDetails(s)
}

func printShapeDetails(shape Shape) {
	fmt.Printf("%s - Area: %.2f\n", shape.Name(), shape.Area())
}
