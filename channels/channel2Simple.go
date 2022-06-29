package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int) //unbuffered channel -
	//znaczy że tylko jedna wiadomość może być w tym
	//samym czasie w channelu

	for j := 0; j < 5; j++ {
		wg.Add(2)

		go func() {
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}()

		go func() {
			ch <- 42
			wg.Done()
		}()
	}

	wg.Wait()
}
