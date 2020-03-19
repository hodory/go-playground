package main

import (
	"fmt"
	"log"

	"github.com/hodory/go-playground/accounts"
)

func main() {
	account := accounts.NewAccount("hodory")
	account.Deposit(100)

	fmt.Println(account.Balance())

	err := account.WithDraw(110)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(account.Balance())
}
