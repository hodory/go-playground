package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	urls := []string{
		"https://www.airbnb.com",
		"https://www.google.co.kr",
		"https://www.naver.com",
	}

	for _, url := range urls {
		hitUrl(url)
	}
}

func hitUrl(url string) error {
	fmt.Println("Checking Url : ", url)
	res, err := http.Get(url)

	if err != nil || res.StatusCode >= 40 {
		return errRequestFailed
	}

	return nil
}
