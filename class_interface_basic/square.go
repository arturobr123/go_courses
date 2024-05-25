// square.go

package main

import (
	"reflect"
)

type Square struct {
	side float64
}

func NewSquare(side float64) *Square {
	return &Square{side}
}

func (s *Square) Area() float64 {
	return s.side * s.side
}

func (c *Square) Name() string {
	return reflect.TypeOf(*c).Name()
}
