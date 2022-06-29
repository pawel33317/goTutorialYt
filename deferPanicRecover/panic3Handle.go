package main

import (
	"fmt"
	"log"
)

//obsługuje wyjątek ale nie wraca na miejsce wystąpienia

func main() {
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
		}
	}()
	panic("something bad happened")
	fmt.Println("end")

}
