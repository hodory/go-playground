package main

import (
	"fmt"

	"std/github.com/hodory/go-playground/banking"
)

func main() {
	account := banking.Account{Owner: "hodory", Balance: 1000}

	fmt.Println(account)
}
