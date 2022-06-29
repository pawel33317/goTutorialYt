package main

import (
	"fmt"
	"time"
)

func main() {
	//gorutyna mimo ża ma inny call stack ma dostęp do stacka głównego
	//jednak są z tym problemy

	var msg = "Hello"
	go func() {
		fmt.Println(msg)
	}()
	msg = "Goodbye" //race condition - lepiej nie używać
	time.Sleep(100 * time.Millisecond)

	var msg2 = "Hello"
	go func(msg2 string) {
		fmt.Println(msg2)
	}(msg2)
	msg2 = "Goodbye" //teraz OK hello bo msg2 jest kopiowana
	time.Sleep(100 * time.Millisecond)
}
