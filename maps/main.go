// A map is a data structure that maps a key-value pair, similar to a dict in Python or C#.
// The keys are statically typed, and all keys must be the same type.
// The values are statically typed, and all values must be the same type.
// The keys and values do not need to be the same type (e.g., an int could be a key for string value).
//
package main

import "fmt"

func main() {
	//colors := map[string]string{
	//	"red":   "#FF0000",
	//	"black": "#000000",
	//	"white": "#FFFFFF",
	//}

	// The zero-value of an unassigned map is an empty map [].
	// var colors map[string]string

	// We can use the make() function to allocate memory and initialize a map.
	colors := make(map[string]string)

	// We can use [] syntax to manually access key-value pairs.
	colors["red"] = "#FF0000"
	colors["white"] = "#FFFFFF"
	colors["black"] = "#000000"

	fmt.Println(colors)

	// We can delete key-value pairs from the map.
	// delete(colors, "white")

	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println("Hex code for", key, "is", value)
	}
}
