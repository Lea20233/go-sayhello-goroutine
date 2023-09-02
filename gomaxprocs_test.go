package go_hello_goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGet_Gomaxprocs(t *testing.T) {

	group := sync.WaitGroup{}
	group.Add(1)

	for i := 0; i < 100; i++ {
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCpus := runtime.NumCPU()
	fmt.Println("Total CPUs = ", totalCpus)

	totalThreads := runtime.GOMAXPROCS(-1)
	fmt.Println("Total threads = ", totalThreads)

	totalGoroutines := runtime.NumGoroutine()
	fmt.Println("Total goroutines = ", totalGoroutines)

}
