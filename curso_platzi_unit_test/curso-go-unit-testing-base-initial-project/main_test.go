package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// func TestMain(t *testing.T) {

// }

func TestAddSuccess(t *testing.T) {
	c := require.New(t)
	result := Add(2, 2)

	expect := 4

	c.Equal(expect, result)
	c.NotEqual(5, result)
}
