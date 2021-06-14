package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)

	go func() { messages <- "Hello from message chanel." }()

	fmt.Println(<-messages)
}
