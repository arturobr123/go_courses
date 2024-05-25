package main

import "fmt"

type shape interface {
	Area() float64
}

type rectangle struct {
	width  float64
	height float64
}

type square struct {
	width  float64
	height float64
}

func (r rectangle) Area() float64 {
	return r.width * r.height
}

func (s square) Area() float64 {
	return s.width * s.height
}

func printArea(s shape) {
	fmt.Println("Area is", s.Area())
}

func main() {
	fmt.Println("Hello, World!")

	rect := rectangle{width: 10, height: 5}
	sq := square{width: 10, height: 10}

	printArea(rect)
	printArea(sq)
}
