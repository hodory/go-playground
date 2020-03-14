package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("Execute After Return")

	length = len(name)
	uppercase = strings.ToUpper(name)

	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func add(numbers ...int) int {
	totalNumber := 0
	for _, number := range numbers {
		totalNumber += number
	}
	return totalNumber
}

func main() {
	totalLength, upper := lenAndUpper("hodory")
	fmt.Println(totalLength, upper)
	fmt.Println(multiply(2, 2))

	repeatMe("react", "vue", "angular", "svelte", "backbone")

	fmt.Println(add(10, 20, 30, 40, 50, 60))
}
