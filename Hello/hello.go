package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Craig")
	fmt.Println(message)
}
