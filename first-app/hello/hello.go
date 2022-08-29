package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names
	names := []string{"Agustin", "Alejandra", "Miqueas"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(messages)
}
