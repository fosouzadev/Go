package main

import (
	"fmt"

	"fosouzadev/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
