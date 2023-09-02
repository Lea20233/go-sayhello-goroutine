package go_hello_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

//using pool in goroutine

func WaitGroup(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
}

func TestPool(t *testing.T) {

	pool := sync.Pool{
		//this is to make default data in pool
		New: func() interface{} {
			return "Default"
		},
	}

	group := &sync.WaitGroup{}

	pool.Put("Hello")
	pool.Put("World")
	pool.Put("This goroutine using pool sync")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			//time.Sleep(1 * time.Second)
			fmt.Println(data)
			pool.Put(data)
		}()
	}
	//time.Sleep(11 * time.Second)
	group.Wait()
	fmt.Println("Done.")
}
