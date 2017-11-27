// package account
// API:
//
// Open(initialDeposit int64) *Account
// (Account) Close() (payout int64, ok bool)
// (Account) Balance() (balance int64, ok bool)
// (Account) Deposit(amount int64) (newBalance int64, ok bool)
//
// If Open is given a negative initial deposit, it must return nil.
// Deposit must handle a negative amount as a withdrawal. Withdrawals must
// not succeed if they result in a negative balance.
// If any Account method is called on an closed account, it must not modify
// the account and must return ok = false.
//
// The tests will execute some operations concurrently. You should strive
// to ensure that operations on the Account leave it in a consistent state.
// For example: multiple goroutines may be depositing and withdrawing money
// simultaneously, two withdrawals occurring concurrently should not be able
// to bring the balance into the negative.
//
// If you are new to concurrent operations in Go it will be worth looking
// at the sync package, specifically Mutexes:
//
// https://golang.org/pkg/sync/
// https://tour.golang.org/concurrency/9
// https://gobyexample.com/mutexes
package account

import (
	"sync"
)

const testVersion = 1

type AccountState int

const (
	open AccountState = iota
	close
)

type Account struct {
	balance int64
	state   AccountState
	sync.RWMutex
}

func Open(initialDeposit int64) (account *Account) {
	if initialDeposit >= 0 {
		account = &Account{balance: initialDeposit, state: open}
	}
	return
}

func (account *Account) Close() (payout int64, ok bool) {
	account.Lock()
	defer account.Unlock()
	if account.state != close {
		account.state = close
		return account.balance, true
	}
	return 0, false
}

func (account *Account) Balance() (balance int64, ok bool) {
	account.RLock()
	defer account.RUnlock()
	if account.state != open {
		return 0, false
	}
	return account.balance, true
}

func (account *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	account.Lock()
	defer account.Unlock()
	if account.state != open {
		return 0, false
	}
	if account.balance+amount >= 0 {
		account.balance += amount
		return account.balance, true
	}
	return account.balance, false
}
