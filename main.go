package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "ramen"}

	john := person{name: "john", age: 20, favFood: favFood}

	fmt.Println(john)
}
