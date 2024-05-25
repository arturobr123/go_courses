package main

import "fmt"

func main() {
	fmt.Println("DICCIONARIOS")

	dictionary := make(map[string]int)

	dictionary["Juan"] = 1
	dictionary["Beto"] = 2

	fmt.Println(dictionary)

	//iterate dictionary
	for key, value := range dictionary {
		fmt.Println(key, value)
	}

	// value exist in dictionary
	value, ok := dictionary["Juan"]
	fmt.Println(value, ok)

	// value does not exist in dictionary
	value2, ok2 := dictionary["ARTURO"]
	fmt.Println(value2, ok2)

}
