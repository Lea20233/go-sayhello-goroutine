package go_hello_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestChannelGoroutine(t *testing.T) {

	//making channel
	channel := make(chan string)

	//this is for close the channel
	defer close(channel)

	//goroutine func
	go func() {
		time.Sleep(1 * time.Second)
		//to send data to channel
		channel <- "Hello World!"
		fmt.Println("'Hello World!' was sent!")
	}()

	//to receice data from channel
	data := <-channel
	fmt.Println(data)

	time.Sleep(2 * time.Second)

}

// this is channel use as parameter
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "This is channel used as parameter!"
}

func TestChannelParameter(t *testing.T) {

	channel := make(chan string)

	defer close(channel)

	go GiveMeResponse(channel)

	result := <-channel
	fmt.Println(result)

}

// channel using in-out only
// channel in for send data
func ChannelIn(channel chan<- string) {
	time.Sleep(1 * time.Second)
	channel <- "This is Channel In-Out"
}

// channel out to receive the ChannelIn
func ChannelOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestChannel_InOut(t *testing.T) {

	channel := make(chan string)

	defer close(channel)

	go ChannelIn(channel)
	go ChannelOut(channel)

	time.Sleep(2 * time.Second)

}

// Buffered Channel = to make capacity data wait in channel
func TestBuffered_Channel(t *testing.T) {

	//this buffer has 3 capacities
	channel := make(chan string, 3)

	defer close(channel)

	go func() {
		channel <- "Hello"
		channel <- "World"
		channel <- "!"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Done.")
}

// use range 'for' in channel if channel have many datas
func TestRange_Channel(t *testing.T) {

	channel := make(chan string)

	go func() {
		for i := 0; i < 12; i++ {
			channel <- "range from " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Receive data", data)
	}

	fmt.Println("Done.")
}

//use 'select' and 'for' if there more than 1 channel to
//receive the datas

func TestSelect_Channel(t *testing.T) {

	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go ChannelIn(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from Channel 1", data)
			counter++
		}
		select {
		case data := <-channel2:
			fmt.Println("Data from Channel 2", data)
			counter++

		}
		//to stop for
		if counter == 2 {
			break
		}
		fmt.Println("Done.")
	}

}

// add default in the select
func TestDefault_SelectChannel(t *testing.T) {

	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go ChannelIn(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data from Channel 2", data)
			counter++
		default:
			fmt.Println("Waiting to load...")

		}
		//to stop for
		if counter == 2 {
			break
		}
	}

}
