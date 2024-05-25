package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	// total := sum(1, 2, 3, 4, 5)

	// if total != 15 {
	// 	t.Errorf("La suma de los valores no es correcta, se esperaba 15 y se obtuvo %d", total)
	// }

	table := []struct {
		values   []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{3, 0, 3, 4, 5}, 15},
		{[]int{0, 0, 3, 7, 5}, 15},
	}

	// tests
	for _, test := range table {
		total := sum(test.values...)
		if total != test.expected {
			t.Errorf("La suma de los valores no es correcta, se esperaba %d y se obtuvo %d", test.expected, total)
		}
	}
}

// go mod init curso_platzi_intermedio/tests
// go test -v
// go test -cover
// go test -coverprofile=coverage.out
// go tool cover -html=coverage.out
