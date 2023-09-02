package go_hello_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer_Time(t *testing.T) {

	// NewTimer on var timer creates a new Timer that will send
	// the current time on timer channel after 5 seconds in var time
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)

}

func TestTime_After(t *testing.T) {

	// time.After waits for 5 seconds to elapse then sends the current time
	// on the channel var time
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <-channel
	fmt.Println(time)

}

func TestAfterFunc_Time(t *testing.T) {

	group := sync.WaitGroup{}
	group.Add(1)

	// AfterFunc waits for 5 seconds to elapse and then calls the current time
	time.AfterFunc(5*time.Second, func() {
		fmt.Println(time.Now())
		group.Done()
	})

	fmt.Println(time.Now())

	group.Wait()

}
