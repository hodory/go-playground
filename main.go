package main

import "fmt"

func main() {
	a := 10
	b := a
	c := &a
	fmt.Println(a, b, c, *c)

	a = 20
	fmt.Println(a, b, c, *c)

	*c = 100
	fmt.Println(a, b, c, *c)
}
