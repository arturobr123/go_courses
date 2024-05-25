package main

import "fmt"

func main() {

	// how defer works
	// defer correra la funcion despues de que se ejecute el programa
	// puede ser para cerrar una conexion a una base de datos, cerrar un archivo, channel, etc
	defer fmt.Println("Hola")

	fmt.Println("Mundo")

	for i := 0; i < 10; i++ {
		if i == 2 {
			fmt.Println("Continue")
			continue
		}

		if i == 8 {
			fmt.Println("Break")
			break
		}

		fmt.Println(i)
	}
}
