package main

import (
	"errors"
	"fmt"
	"net/http"
)

// Note: This is checking each link synchronously!
// Each link is checked in turn, waiting for the previous to finish before starting the next.
func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://example.com",
	}

	checkLinksSync(links)
}

func checkLinksSync(links []string) {
	for _, link := range links {
		err := checkLinkSync(link)

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Link appears up! %v\n", link)
		}
	}
}

func checkLinkSync(link string) error {
	resp, err := http.Get(link)

	if err != nil {
		return err
	}

	if resp.StatusCode == 200 {
		return nil
	} else {
		return errors.New("Link might be down! " + link)
	}
}
