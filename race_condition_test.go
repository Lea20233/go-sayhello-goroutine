package go_hello_goroutine

import (
	"fmt"
	"testing"
	"time"
)

//testing race condition with multiple access to variable X
//where the test results are different every it runs

func TestRace_Condition(t *testing.T) {

	x := 0

	for i := 1; i <= 10000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)

}
