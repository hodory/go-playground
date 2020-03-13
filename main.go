package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)

	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	totalLength, upper := lenAndUpper("hodory")
	fmt.Println(totalLength, upper)
	fmt.Println(multiply(2, 2))

	repeatMe("react", "vue", "angular", "svelte", "backbone")
}
