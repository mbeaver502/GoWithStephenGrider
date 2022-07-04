package main

import "fmt"

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

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}
