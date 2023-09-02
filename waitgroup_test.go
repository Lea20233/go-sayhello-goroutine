package go_hello_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchron(group *sync.WaitGroup) {

	defer group.Done()

	group.Add(1)

	fmt.Println("Hello World")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {

	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go RunAsynchron(group)
	}

	group.Wait()
	fmt.Println("Done")

}
