package main

import (
	"errors"
	"fmt"
)

// IProduct defines the interface for products created by the factory.
type IProduct interface {
	getName() string
	setName(name string)
	getStock() int
	setStock(stock int)
}

// Computer is a struct that represents a generic computer with a name and stock level.
type Computer struct {
	name  string
	stock int
}

func (c *Computer) getName() string {
	return c.name
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getStock() int {
	return c.stock
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}

// Laptop is a struct that embeds Computer, representing a specific type of Computer.
type Laptop struct {
	Computer
}

// newLaptop creates a new instance of Laptop, initializing it with predefined values.
func newLaptop() IProduct {
	return &Laptop{Computer{name: "Laptop", stock: 25}}
}

// Desktop is another struct that embeds Computer, representing another specific type of Computer.
type Desktop struct {
	Computer
}

// newDesktop creates a new instance of Desktop, initializing it with predefined values.
func newDesktop() IProduct {
	return &Desktop{Computer{name: "Desktop", stock: 15}}
}

// getComputerFactory is a factory function that returns a new instance of a product based on the input type.
func getComputerFactory(computerType string) (IProduct, error) {
	switch computerType {
	case "laptop":
		return newLaptop(), nil
	case "desktop":
		return newDesktop(), nil
	default:
		return nil, errors.New("Invalid computer type")
	}
}

// printNameAndStock prints the name and stock of a product.
func printNameAndStock(product IProduct) {
	fmt.Printf("Name: %s, Stock: %d\n", product.getName(), product.getStock())
}

func main() {
	fmt.Println("FACTORY PATTERN DEMONSTRATION")

	// Creating a laptop using the factory function
	laptop, _ := getComputerFactory("laptop")
	// Creating a desktop using the factory function
	desktop, _ := getComputerFactory("desktop")

	// Displaying the details of the products
	printNameAndStock(laptop)
	printNameAndStock(desktop)
}
