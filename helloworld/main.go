/*
	We can execute some commands from terminal:

		go run <filenames>		Build *and* execute one or more Go files
		go build <filenames> 	Build one or more Go files

*/

// In Go, a package is like a project or workspace or namespace
// A package can have many related Go files inside it -- but each must declare the same package name
// Two types of packages:
//		(1) Executable -- that which generates an executable file
//		(2) Reusable -- that which can be re-used as "helper" methods (e.g., code libraries)
// How do we know if a package is executable or reusable? The name of the package determines this!
// 		The package name of "main" tells the Go compiler that this is an executable package!
//		An executable package *must* have a func main!
// 		Any other name will tell Go that this is *not* an executable package!
package main

// The fmt package gives us access to std I/O operations like Println
import "fmt"

// Since this is an executable package (package main), we *must* have a func main entrypoint
func main() {
	fmt.Println("hello world :)")
}
