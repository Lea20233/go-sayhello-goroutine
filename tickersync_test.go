package go_hello_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker_Sync(t *testing.T) {

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	//make channel 'done'
	done := make(chan bool)

	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()

	for {
		select {
		case <-done:
			fmt.Println("Done.")
			return

		case x := <-ticker.C:
			fmt.Println("Time = ", x)
		}
	}
}

func TestTick(t *testing.T) {

	channel := time.Tick(1 * time.Second)

	for tick := range channel {
		fmt.Println(tick)
	}

}
