package main

import (
	"fmt"
	"sync"
)

//pokazanie skąd go wie, że channel jest już zamknięty
var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50)
	wg.Add(2)

	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}

//nie można wysyłać na close channel bo panic
//nie można otworzyć ponownie channela
