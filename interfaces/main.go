package main

import "fmt"

// To whom it may concern, here's a contract that anyone can use inside this program.
// Any variable of type "bot" can perform these actions:
//  (1) getGreeting() returns a string
// Any type in this program that has a getGreeting() string function is now an honorary "bot".
// So anywhere there's a type "bot" being used, you can use that more concrete type (e.g., englishBot).
// It's like we have a public class EnglishBot : IBot , where a class implements an interface.
type bot interface {
	// functionName(argTypes) returnTypes
	getGreeting() string
}

// Empty structs for the sake of example.
// Imagine these structs have similar functionality but different implementation details.
type englishBot struct{}
type spanishBot struct{}

func (eb englishBot) getGreeting() string {
	// Some complicated logic that's different to the spanishBot version of getGreeting...
	return "Hello! :)"
}

func (sb spanishBot) getGreeting() string {
	// Some complicated logic that's different to the englishBot version of getGreeting...
	return "Hola! :)"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}
