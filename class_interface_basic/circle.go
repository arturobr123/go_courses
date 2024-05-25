// circle.go

package main

import (
	"math"
	"reflect"
)

type Circle struct {
	radius float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{radius}
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Name() string {
	return reflect.TypeOf(*c).Name()
}
