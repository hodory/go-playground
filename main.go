package main

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	totalLength, _ := lenAndUpper("hodory")
	fmt.Println(totalLength)
	fmt.Println(multiply(2, 2))

	repeatMe("react", "vue", "angular", "svelte", "backbone")
}
