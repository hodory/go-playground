package main

import (
	"fmt"

	"github.com/hodory/go-playground/accounts"
)

func main() {
	account := accounts.NewAccount("hodory")

	fmt.Println(account)
}
