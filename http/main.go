package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// See: https://pkg.go.dev/net/http@go1.18.3#Get
	resp, err := http.Get("https://www.example.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// The response body (resp.Body) is of the type io.ReadCloser, which itself an interface!
	// Multiple interfaces can be combined to create a new interface.
	// ReadCloser is an interface that combines the Reader interface and the Closer interface.
	// The Reader interface allows us to use the same functionality across many possible data sources.
	// Note that body comes back as a slice of byte, which is pretty flexible.
	body, err := io.ReadAll(resp.Body)

	// We *MUST* close the response body.
	// We get access to Close() because resp.Body implements io.ReadCloser.
	// io.ReadCloser implements Closer, which gives us Close().
	resp.Body.Close()

	if err != nil {
		fmt.Println("Error:", err)
	}

	if resp.StatusCode == 200 {
		// Cast the body (which is a slice of byte) to string for easier readability.
		fmt.Println(string(body))
	} else {
		fmt.Println(resp.Status)
	}
}
