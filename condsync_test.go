package go_hello_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var condition = sync.NewCond(&sync.Mutex{})
var group = &sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	condition.L.Lock()
	fmt.Println("Done", value)
	condition.L.Unlock()
}

func TestCondition_Wait(t *testing.T) {

	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			//Signal wakes one goroutine waiting on i, if there is any
			condition.Signal()
		}
	}()

	//go func() {
	//	time.Sleep(1 * time.Second)
	//Broadcast wakes all goroutines waiting on i
	//	condition.Broadcast()
	//}()

	group.Wait()
}
