package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	var results = map[string]string{}

	urls := []string{
		"https://www.airbnb.com",
		"https://www.google.co.kr",
		"https://www.naver.com",
	}

	for _, url := range urls {
		result := "OK"
		err := hitUrl(url)

		if err != nil {
			result = "FAILED"
		}

		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitUrl(url string) error {
	fmt.Println("Checking Url : ", url)
	res, err := http.Get(url)

	if err != nil || res.StatusCode >= 400 {
		return errRequestFailed
	}

	return nil
}
