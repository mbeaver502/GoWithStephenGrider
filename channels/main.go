package main

import (
	"fmt"
	"net/http"
	"time"
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

	// We've created a go-routine for each link. So we should be able to iterate that many times and wait.
	// Waiting for the channel message is still blocking, but now we're reasonably guaranteed to get messages from each go-routine.
	// for i := 0; i < len(links); i++ {
	//
	// Infinite loop to keep pinging links forever.
	// We wait to get a response back from the child go-routine and then launch another for the same link.
	// We are communicating strings over our channel, so we can use the message from that channel as an arg for checkLink().
	// for {
	// 	go checkLink(<-c, c)
	// }
	//
	// Alternative for-loop syntax.
	// We can use range on a channel to wait for a channel to return a message.
	// The message from our channel gets assigned to the variable l, which we can then use as an arg for checkLink().
	// This is equivalent to the infinite loop above, but this is clearer in purpose.
	for l := range c {
		// This will launch a new go-routine as soon as possible.
		// So we're going to end up spamming each link as we spawn new go-routines as soon as the previous ends.
		// go checkLink(l, c)

		// Create and call a function literal (lambda function).
		// We're using a lambda so that we can sleep before the next go-routine is spawned.
		// We do this so that we're not unintentionally blocking our main go-routine with the sleep.
		// The newly spawned lambda will first sleep and then execute checkLink(). Notice that checkLink() is not itself spawned as a go-routine.
		// We sleep after each go-routine ends so that we're not spamming the links.
		// After we've slept for an adequate time, then we can launch a new go-routine for the recently finished link.
		// Recall that we're receiving the previously finished link over the channel.
		// Since the link we get back from the channel can change at any time, we want to make sure we use the right one.
		// We can ensure we use the right link (l) by passing it as an arg to our lambda function.
		go func(link string) {
			time.Sleep(3 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	// http.Get() is a blocking call, so our routine will be forced to wait until it completes.
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
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
