/*
	We can execute some commands from terminal:

		go run <filenames>		Build *and* execute one or more Go files
		go build <filenames> 	Build one or more Go files

*/
package main

// the fmt package gives us access to std I/O operations like Println
import "fmt"

/*
	This is a multi-line comment.
*/
func main() {
	fmt.Println("hello world :)")
}
