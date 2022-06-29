package main

import (
	"fmt"
	"sync"
)

//pozwalają na bezpieczne przesyłanie danych między go rutynami

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int) //int będzie przesyłany przez channel
	wg.Add(2)

	go func() { //receiver
		i := <-ch //otrzymuje dane z channelu
		fmt.Println(i)
		wg.Done()
	}()

	go func() {
		ch <- 42 //wrzuca do channelu - KOPIĘ
		wg.Done()
	}()

	wg.Wait()
}
