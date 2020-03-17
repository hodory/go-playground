package main

import "fmt"

func main() {
	data := map[string]int{"age": 12, "height": 180, "weight": 70}

	for key, value := range data {
		fmt.Println(key, value)
	}
}
