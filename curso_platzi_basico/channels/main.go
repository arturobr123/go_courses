package main

import "fmt"

func say(text string, c chan<- string) {
	c <- text
}

func main() {
	c := make(chan string, 2)

	fmt.Println("FIRST")

	go say("Hello1", c)
	go say("Hello2", c)

	// msg := <-c //obtiene el valor de el channel
	// fmt.Println(msg)

	// msg2 := <-c //obtiene el valor de el channel
	// fmt.Println(msg2)

	// WAIT FOR ALL EVENTS TO BE COMPLETED IN CHANNEL
	for i := 0; i < cap(c); i++ {
		msg := <-c //obtiene el valor de el channel
		fmt.Println(msg)
	}
}

// QUESTION
// Question, both events in channel are running at the same time ?
// and in for we are waiting for all events to be completed ?

// ANSWER
// Yes, in your Go code, both say("Hello1", c) and say("Hello2", c) are launched concurrently
// using goroutines (go keyword). This means they are running in parallel, and each can execute independently.
// The channel c is used to communicate between these goroutines and the main goroutine.
// Since the channel has a buffer size of 2, it can hold both messages without blocking the sender goroutines until the buffer is full.
// In the for loop, you are indeed waiting for all events to be completed, or more specifically, for all messages to be received from the channel.
// The loop iterates twice (i < cap(c)) because the capacity of the channel c is 2.
// Each iteration of the loop retrieves one message from the channel with msg := <-c and prints it.
// This loop ensures that the main goroutine waits and processes all messages sent to the channel before it exits.
