package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (my_pc pc) ping() {
	fmt.Println(my_pc.brand, "is pinging")
}

// como agrego *pc para acceder a los atributos de la estructura por medio de su direccion de memoria
// asi se accede a la memoria y se actualiza de manera mas optima
func (my_pc *pc) duplicateRam() {
	my_pc.ram = my_pc.ram * 2
}

func main() {
	fmt.Println("Hello, World!")

	a := 10
	b := &a // b is a pointer to a , direccion de memoria de a

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b) // *b is the value of b, the value of a

	*b = 20
	fmt.Println(a) // es el misma direccion de memoria entonces si modifico el valor de b, modifico el valor de a

	//////////////////////////////////////////////////////////////////////////////////////////////

	my_pc := pc{ram: 8, disk: 100, brand: "Dell"}
	fmt.Println(my_pc)
	my_pc.ping()

	my_pc.duplicateRam()
	fmt.Println(my_pc)

	my_pc.duplicateRam()
	fmt.Println(my_pc)

}
