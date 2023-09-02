package go_hello_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//using mutex to fix previous race condition
//to lock and unlock goroutine who can access variable X

func TestMutex(t *testing.T) {

	x := 0

	var mutex sync.Mutex

	for i := 1; i <= 10000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)

}

//this is RWMutex

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestAccount(t *testing.T) {
	acc := BankAccount{}

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 5; j++ {
				acc.AddBalance(1)
				fmt.Println(acc.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance In Your Account $", acc.GetBalance())
}
