package go_hello_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnceSync() {
	counter++

}

func TestSync_Once(t *testing.T) {

	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnceSync)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter", counter)

}
