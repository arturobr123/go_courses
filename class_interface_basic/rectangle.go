// rectangle.go

package main

import (
	"reflect"
)

type Rectangle struct {
	length, width float64
}

func NewRectangle(length, width float64) *Rectangle {
	// The '&' operator is used to get the address of a new Rectangle instance.
	// Returning '*Rectangle' means this function returns a pointer to a Rectangle instance.
	return &Rectangle{length, width}
}

func (r *Rectangle) Area() float64 {
	return r.length * r.width
}

func (c *Rectangle) Name() string {
	return reflect.TypeOf(*c).Name()
}
