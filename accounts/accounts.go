package accounts

import "errors"

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNotEnough = errors.New("You don't have enough balance")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// WithDraw x amount from your account
func (a *Account) WithDraw(amount int) error {
	if a.balance < amount {
		return errNotEnough
	}
	a.balance -= amount
	return nil
}

//ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}
