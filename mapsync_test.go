package go_hello_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func AddDataMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()
	group.Add(1)

	//to add/save data in map
	data.Store(value, value)

}

func TestMap_Sync(t *testing.T) {

	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		go AddDataMap(data, i, group)
	}

	group.Wait()

	//to iterate data
	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})

}
