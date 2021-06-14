package main

import (
	"fmt"
	"time"
)

func print(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ": ", i)
	}
}

func main() {
	print("direct")

	go print("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
