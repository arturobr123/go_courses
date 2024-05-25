package main

import (
	"fmt"
	"time"
)

type Person struct {
	DNI  string
	Name string
	Age  int
}

type Employee struct {
	ID       int
	Position string
}

type FullTimeEmployee struct {
	Employee
	Person
}

var GetPersonByDNI = func(dni string) (Person, error) {
	_ = dni // explicitly ignore the parameter

	time.Sleep(3 * time.Second)
	return Person{}, nil
}

var GetEmployeeByID = func(id int) (Employee, error) {
	_ = id // explicitly ignore the parameter

	time.Sleep(3 * time.Second)
	return Employee{}, nil
}

func GetFullTimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	// PERSON
	var fullTimeEmployee FullTimeEmployee
	person, err := GetPersonByDNI(dni)
	if err != nil {
		return fullTimeEmployee, err
	}
	fullTimeEmployee.Person = person

	// EMPLOYEE
	employee, err := GetEmployeeByID(id)
	if err != nil {
		return fullTimeEmployee, err
	}
	fullTimeEmployee.Employee = employee

	return fullTimeEmployee, nil
}

func main() {
	fmt.Println("Hello, World!")
}
