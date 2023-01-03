package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	// A slice of sample websites
	urls := []string{
		"https://www.facebook.com/",
		"https://www.face.com/",
		"https://www.google.com/",
		"https://www.twitter.com/",
	}
	var wg sync.WaitGroup

	for _, u := range urls {
		// Increment the wait group counter
		wg.Add(1)
		go func(url string) {
			// Decrement the counter when the go routine completes
			defer wg.Done()
			// Call the function check
			checkUrl(url)
		}(u)
	}
	// Wait for all the checkWebsite calls to finish
	wg.Wait()
}

// checks and prints a message if a website is up or down
func checkUrl(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is down !!!")
		return
	}
	fmt.Println(url, "is up and running.")
}
