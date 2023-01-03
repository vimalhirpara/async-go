package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// A slice of sample websites
	urls := []string{
		"https://www.facebook.com/",
		"https://www.face.com/",
		"https://www.google.com/",
		"https://www.twitter.de/",
	}
	for _, url := range urls {
		go checkUrl(url)
	}
	time.Sleep(5 * time.Second)
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
