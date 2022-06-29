package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50)
	wg.Add(2)

	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		} //deadlock teraz tutaj bo nie ma końca
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		wg.Done()
	}(ch)

	wg.Wait()
}
