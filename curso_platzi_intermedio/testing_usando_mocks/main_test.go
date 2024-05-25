package main

import (
	"testing"
)

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id                       int
		dni                      string
		mockFunc                 func()
		expectedFullTimeEmployee FullTimeEmployee
	}{
		{1, "1", func() {
			GetPersonByDNI = func(dni string) (Person, error) {
				return Person{
					Name: "Arturo",
					Age:  27,
					DNI:  "1",
				}, nil
			}
			GetEmployeeByID = func(id int) (Employee, error) {
				return Employee{
					ID:       1,
					Position: "CEO",
				}, nil
			}
		}, FullTimeEmployee{
			Employee: Employee{
				ID:       1,
				Position: "CEO",
			},
			Person: Person{
				Name: "Arturo",
				Age:  27,
				DNI:  "1",
			},
		}},
	}

	//use table
	for _, test := range table {
		test.mockFunc()
		fullTimeEmployee, err := GetFullTimeEmployeeById(test.id, test.dni)
		if err != nil {
			t.Errorf("Error getting full time employee by id: %v", err)
		}
		if fullTimeEmployee != test.expectedFullTimeEmployee {
			t.Errorf("Expected full time employee: %v, got: %v", test.expectedFullTimeEmployee, fullTimeEmployee)
		}
	}
}

// go mod init github.com/arturo-c/curso_platzi_intermedio/testing_usando_mocks
// go test -v
