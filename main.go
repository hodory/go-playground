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

func print(numbers ...int) int {
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i])
	}
	return 1
}

func main() {
	totalLength, upper := lenAndUpper("hodory")
	fmt.Println(totalLength, upper)
	fmt.Println(multiply(2, 2))

	repeatMe("react", "vue", "angular", "svelte", "backbone")

	print(1, 2, 3, 4, 5, 6)
}
