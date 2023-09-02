package go_hello_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func TransferBalance(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Locking User 1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Locking user 2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()

}

func TestTransfer_Balance(t *testing.T) {

	user1 := UserBalance{
		Name:    "Luna",
		Balance: 5000,
	}

	user2 := UserBalance{
		Name:    "Nana",
		Balance: 5000,
	}

	go TransferBalance(&user1, &user2, 1000)
	go TransferBalance(&user2, &user1, 2000)

	time.Sleep(5 * time.Second)

	fmt.Println("Account Balance for", user1.Name, "=", user1.Balance)
	fmt.Println("Account Balance for", user2.Name, "=", user2.Balance)

}
