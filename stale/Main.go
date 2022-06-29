package main

import (
	"fmt"
)

const (
	b = iota //iota to counter
	n = iota
	m = iota
)

const (
	b2 = iota //iota to counter
	n2
	m2
)

const (
	_  = iota //odrzuca
	b3 = iota //iota to counter
)

const (
	_  = iota + 6 //odrzuca i ustawia start
	b4            //iota to counter
)

func main() {
	const myConst int = 42 //nazewnictwo takie ssamo jak zwykłych
	fmt.Printf("%v, %T\n", myConst, myConst)

	//nie można przypisać wyniku funkcji bo to jest real time robione
	//const nextConst float64 = math.Sin(1.57)
	//fmt.Printf("%v, %T", nextConst, nextConst)

	const constA = 42
	fmt.Printf("%v, %T\n", constA, constA)

	const constB int16 = 34
	fmt.Printf("%v, %T\n", constA+constB, constA+constB) //zadziałą implicit konwersja

	fmt.Println(b, " ", n, " ", m, " ")
	fmt.Println(b2, " ", n2, " ", m2, " ")
	fmt.Println(b3)
	fmt.Println(b4)

	//stałe muszą być znane w tompiletime

	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
	)
	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)

	const (
		isAdmin = 1 << iota
		isMod
		isUser
		flag4
		flag5
	)

	var roles byte = isAdmin | isUser | flag5
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v", isAdmin&roles == isAdmin)
}
