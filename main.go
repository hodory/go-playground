package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	people := [5]string{"foo", "bar", "baz", "buz", "joo"}

	for _, person := range people {
		go isNice(person, c)
	}

	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
}

func isNice(name string, c chan bool) {
	time.Sleep(time.Second * 5)
	c <- true
}
