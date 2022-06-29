package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)

	wg.Add(2)

	go func(ch <-chan int) { //receive only channel
		//castuje channel na jednokierunkowy, rzadkie w go
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	go func(ch chan<- int) { //sending to the channel only
		ch <- 42
		//fmt.Println(<-ch) error bo tylko można pisać do channela
		wg.Done()
	}(ch)

	wg.Wait()
}
