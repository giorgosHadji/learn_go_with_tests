package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Technically you do not need to change Balance to use a pointer receiver
// as taking a copy of the balance is fine. However, by convention you should keep your method receiver types the same for consistency.
func (w *Wallet) Deposit(initialBalance Bitcoin) {
	// panic("unimplemented")s
	w.balance += initialBalance
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)

}

// The var keyword allows us to define values global to the package.

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance > amount {
		w.balance -= amount
		return nil
	}
	return ErrInsufficientFunds

}
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

/*
Pointers
Go copies values when you pass them to functions/methods, so if you're writing a function that needs to mutate state you'll need it to take a pointer to the thing you want to change.

The fact that Go takes a copy of values is useful a lot of the time but sometimes you won't want your system to make a copy of something, in which case you need to pass a reference. Examples include referencing very large data structures or things where only one instance is necessary (like database connection pools).

nil
Pointers can be nil

When a function returns a pointer to something, you need to make sure you check if it's nil or you might raise a runtime exception - the compiler won't help you here.

Useful for when you want to describe a value that could be missing

Errors
Errors are the way to signify failure when calling a function/method.

By listening to our tests we concluded that checking for a string in an error would result in a flaky test. So we refactored our implementation to use a meaningful value instead and this resulted in easier to test code and concluded this would be easier for users of our API too.

This is not the end of the story with error handling, you can do more sophisticated things but this is just an intro. Later sections will cover more strategies.
*/
