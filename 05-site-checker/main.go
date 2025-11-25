package main

import (
	"fmt"
	"net/http"
	"time"
)

var urls = []string{
	"https://www.google.com",
	"https://www.facebook.com",
	"https://www.stackoverflow.com",
	"https://www.go.dev",
	"https://www.amazon.com",
	"https://www.reddit.com",
	"https://www.microsoft.com",
	"https://www.wikipedia.org",
}

type CheckResult struct {
	URL    string
	Status string
	IsUp   bool
}

func main() {
	start := time.Now()

	c := make(chan CheckResult)

	fmt.Println("Start checking websites with Channels...")

	for _, url := range urls {
		go checkUrl(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c 

		if result.IsUp {
			fmt.Printf("[UP]   %s \n", result.URL)
		} else {
			fmt.Printf("[DOWN] %s (Status: %s)\n", result.URL, result.Status)
		}
	}

	fmt.Printf("\nTotal time taken: %v\n", time.Since(start))
}

func checkUrl(url string, c chan<- CheckResult) {
	resp, err := http.Get(url)
	
	status := "OK"
	isUp := true

	if err != nil {
		status = fmt.Sprintf("Error: %v", err)
		isUp = false
	} else {
		defer resp.Body.Close()
		status = fmt.Sprintf("%d", resp.StatusCode)
		if resp.StatusCode != 200 {
			isUp = false
		}
	}

	c <- CheckResult{
		URL:    url,
		Status: status,
		IsUp:   isUp,
	}
}
