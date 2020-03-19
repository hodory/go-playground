package main

import (
	"fmt"

	"github.com/hodory/go-playground/accounts"
)

func main() {
	account := accounts.NewAccount("hodory")
	account.Deposit(100)

	fmt.Println(account.Balance())

	account.WithDraw(30)
	fmt.Println(account.Balance())
}
