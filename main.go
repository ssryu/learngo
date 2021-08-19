package main

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	errRequestFailed = errors.New("Request failed")
)

type requestResult struct {
	url    string
	status string
}

func main() {
	c := make(chan requestResult)
	results := make(map[string]string)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- requestResult) {
	resp, _ := http.Get(url)
	c <- requestResult{url: url, status: resp.Status}
}
