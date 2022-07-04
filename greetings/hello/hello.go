package main

import (
	"fmt"
	"log"
	"mbeaver/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//message, err := greetings.Hello("Gladys")

	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
