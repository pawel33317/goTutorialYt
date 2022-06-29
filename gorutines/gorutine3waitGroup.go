package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	var msg2 = "Hello"
	wg.Add(1)
	go func(msg2 string) {
		fmt.Println(msg2)
		wg.Done()
	}(msg2)

	wg.Wait()
}
