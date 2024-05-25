package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEployee struct {
	Person
	Employee
}

func main() {
	ftEmployee := FullTimeEployee{}
	ftEmployee.id = 1
	ftEmployee.name = "Maria"
	ftEmployee.age = 27
	fmt.Printf("%v", ftEmployee)
}
