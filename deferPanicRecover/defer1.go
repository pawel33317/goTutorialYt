package main

import "fmt"

//Defer - wywołanie funkcji w przyszłości
//Panic - kiedy aplikacja pada
//Recover - naprawa panic

func main() {
	fmt.Println("start")
	defer fmt.Println("middle") //wywołanie po wszystkich funkcjach ale przed return
	fmt.Println("end")          //LIFO jak będzie więcej defer funkcji

	a := "a start"
	defer fmt.Println(a) //Defer kopiuje wartość do momentu wywołania
	a = "a end"
}
