package main

import "fmt"

func main() {
	names := [5]string{"john", "jay", "mac", "cindy"}
	names[4] = "david"

	fmt.Println(names)
}
