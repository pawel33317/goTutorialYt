package main

import (
	"fmt"
	"log"
)

//obsługuje wyjątek ale nie wraca na miejsce wystąpienia

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
		}
	}()
	panic("something bad")
	fmt.Println("done panicking")
}
