package main

import (
	"fmt"
	"strconv"
)

//na zewnątrz trzeba zrobić pełną deklarację
var globalVar float32 = 4.

//deklaracja grupy zmiennych
var (
	pName string = "Pawel"
	pAge  int    = 30
)

//zmienne zadeklarowane z małej litery są lokalne dla package a z dużej exportowane
var I int = 42

func main() {
	var i int
	i = 42
	fmt.Println(i)

	var i2 int = 42
	fmt.Println(i2)

	i3 := 42
	fmt.Println(i3)

	i4 := 4.0
	fmt.Printf("%v, %T\n", i4, i4) //jest też float32 ale trzeba ręcznie zadeklarować

	fmt.Printf("%v, %T\n", globalVar, globalVar)

	fmt.Printf("%v[%T], %v[%T]\n", pName, pName, pAge, pAge)

	//akronimy z dużej np: theHTTPRequest

	//konwersja
	var i5 float64 = float64(i)
	fmt.Printf("%v, %T\n", i5, i5)

	var i6 string
	i6 = strconv.Itoa(i3)
	fmt.Printf("%v, %T\n", i6, i6)
}
