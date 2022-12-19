package main

import (
	"fmt"

	"my_test_job/greetings/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
