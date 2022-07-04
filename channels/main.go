package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://example.com",
	}

	// Note: This is checking each link synchronously!
	// Each link is checked in turn, waiting for the previous to finish before starting the next.
	// checkLinksSync(links)

	for _, link := range links {
		// Launch a new go-routine for checkLink().
		//
		// When we launch multiple go-routines, the Go scheduler handles them.
		// By default, Go attempts to use only one CPU core. We can change that so Go uses multiple CPU cores.
		// If we have one CPU core, then the Go Scheduler runs *one* routine until it finishes or makes a blocking call (like http.Get()).
		// If Go is working on multiple CPU cores, then the Go Scheduler will run one thread per each "logical" core.
		//
		// Concurrency is not parallelism!
		// Concurrency -- We can have multiple threads executing code. If one thread blocks, another one is picked up and worked on.
		// Parallelism -- Multiple threads executed at the *exact same time*. This requires multiple CPU cores.
		//
		go checkLink(link)
	}
}

func checkLink(link string) {
	// http.Get() is a blocking call, so our routine will be forced to wait until it completes.
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
}

/*
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
*/
