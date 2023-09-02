package go_hello_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World! Nice to meet you!")
}

func TestGoroutine(t *testing.T) {
	//this 'go' is goroutine for func RunHelloWorld
	go RunHelloWorld()
	fmt.Println("Hi!")

	//time.Sleep use for giving the goroutine time for execute all the codes
	//so the program won't end abruptly with code unexecute
	time.Sleep(1 * time.Second)

}

// another example using goroutine
func DisplayNumbers(number int) {

	fmt.Println("Display number ", number)

}

// use goroutines for 'for'
func TestGoroutines(t *testing.T) {

	for i := 0; i < 10000; i++ {
		go DisplayNumbers(i)
	}

	time.Sleep(10 * time.Second)

}
