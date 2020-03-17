package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "ramen"}

	john := person{"john", 20, favFood}

	fmt.Println(john)
	// for key, value := range john {
	// 	fmt.Println(key, value)
	// }
}
