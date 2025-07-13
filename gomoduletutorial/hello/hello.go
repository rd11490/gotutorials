package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Ryan", "Minnie", "Leon", "Evie"}

	message, err := greetings.Hello("Ryan")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
