package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("PrinceCharming")
	fmt.Println(message)
}
