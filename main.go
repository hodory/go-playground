package main

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func main() {
	results := map[string]string{}
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com",
		"https://www.google.co.kr",
		"https://www.naver.com",
	}

	for _, url := range urls {
		go hitUrl(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitUrl(url string, c chan<- requestResult) {
	fmt.Println("Checking Url : ", url)
	res, err := http.Get(url)
	status := "OK"
	if err != nil || res.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}
