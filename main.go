package main

import "fmt"

func main() {
	names := []string{"john", "jay", "mac", "cindy"}
	names = append(names, "tommy")

	fmt.Println(names)
}
