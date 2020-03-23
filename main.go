package main

import (
	"fmt"
	"time"
)

func main() {
	go printText("Hello")
	printText("World")
}

func printText(text string) {
	for i := 0; i < 10; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
