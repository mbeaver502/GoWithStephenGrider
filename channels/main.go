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

	// We can use channels to communicate between the main go-routine and child go-routines.
	// Make a new channel that can be communicated over using string type.
	c := make(chan string)

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
		go checkLink(link, c)
	}

	// Waiting for a message to come back over the channel is a blocking call.
	// Once we get something back over the channel, we immediately print it out and the main go-routine exits.
	// Receiving channel messages is a blocking action!
	// fmt.Println(<-c)

	// We've create a go-routine for each link. So we should be able to iterate that many times and wait.
	// Waiting for the channel message is still blocking, but now we're reasonably guaranteed to get messages from each go-routine.
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	// http.Get() is a blocking call, so our routine will be forced to wait until it completes.
	_, err := http.Get(link)

	if err != nil {
		c <- link + " might be down!"
		// fmt.Println(link, "might be down!")
		return
	}

	c <- link + " is up!"
	// fmt.Println(link, "is up!")
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
