package main

import (
	"fmt"

	basemod "github.com/abhishekunotech/go-module-base"
)

func main() {
	// Get a greeting message and print it.
	message := basemod.Hello("Gladys")
	fmt.Println(message)
}
