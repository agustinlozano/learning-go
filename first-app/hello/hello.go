package main

import (
	"fmt"
	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Agustin")
	fmt.Println(message)
}
