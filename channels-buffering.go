package main

import (
	"fmt"
)

func main() {
	// Here we make a channel of buffering up to 2 messages
	// If do not do this, we will get
	// "fatal error: all goroutines are asleep - deadlock!" error
	messages := make(chan string, 2)

	messages <- "Message 1\r"
	messages <- "Message 2\r"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
