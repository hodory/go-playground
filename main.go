package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	people := [3]string{"foo", "bar", "baz"}

	for _, person := range people {
		go isNice(person, c)
	}

	result := <-c
	fmt.Println(result)
}

func isNice(name string, c chan bool) {
	time.Sleep(time.Second * 5)
	c <- true
}
