package main

import "fmt"

func message(msg string, c chan string) {
	c <- msg
}

func main() {
	fmt.Println("Hello, World!")

	email1 := make(chan string)
	email2 := make(chan string)

	go message("Hello1", email1)
	go message("Hello2", email2)

	for i := 0; i < 2; i++ {
		select {
		case msg := <-email1:
			fmt.Println("Email recibido de email1: ", msg)
		case msg2 := <-email2:
			fmt.Println("Email recibido de email2: ", msg2)
		}
	}

	close(email1) // Closing the channel email1 to signal that no more data will be sent on this channel.
	close(email2) // Closing the channel email2 serves the same purpose, ensuring that receivers can terminate gracefully.
}
