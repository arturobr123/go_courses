package main

import (
	"fmt"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done() // decrement the wait group by 1
	fmt.Println(text)
}

func main() {

	var wg sync.WaitGroup // acumula un conjunto de go routines y las va liberando

	fmt.Println("Hello")
	wg.Add(1) // add 1 to the wait group

	go say("World", &wg) // concurrency, running this function in the background

	wg.Wait() // wait for the wait group to be done, espera a que todas las go routines hayan terminado

	// FUNCIONES ANONIMAS - igual con concurrency usando goroutines
	// go func(text string) {
	// 	fmt.Println(text)
	// }("Hello")

	//time.Sleep(time.Second * 1)
}
