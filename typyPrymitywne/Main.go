package main

import "fmt"

//bool
//integer, floating point, complex
//text

func main() {
	var n bool = true
	fmt.Printf("%v[%T]\n", n, n)

	//int minimum 32bity
	//domyślnie inicjalizowane zerem
	x1 := int8(99) //int16, int32, int64, uint8, uint16, uint32,   byte == uint8
	fmt.Printf("%v[%T]\n", x1, x1)

	x2 := x1 * x1
	fmt.Printf("%v[%T]\n", x2, x2)

	//nie można dodawać różnych typów

	//operatory bitowe, &, |, ^, &^
	//bitshift << >>

	//float32, float64

	var c complex128 = 2 + 4i //jest jeszcze complex64
	var c2 complex128 = complex(6, 7)
	fmt.Printf("%v[%T]\n", c, c)
	fmt.Printf("%v[%T]\n", real(c), real(c))
	fmt.Printf("%v[%T]\n", imag(c2), imag(c2))

	//string  w go to dowolny utf 8 charakter - więc nie można zakodować wielu znaków
	// można traktować jak tablice
	// stringi są niezmienne więc nie można podmienić znaku

	s := "this is stringg"
	s2 := "string2"
	fmt.Printf("%v[%T]\n", s, s)
	fmt.Printf("%v[%T]\n", s[0], s[0])
	fmt.Printf("%v[%T]\n", s+s2, s+s2)

	kolekcjaBajtow := []byte(s)
	fmt.Printf("%v[%T]\n", kolekcjaBajtow, kolekcjaBajtow)
	//na kolekcji bajtów już można operować :)

	s4 := "Hello, 世界" //też string why???
	fmt.Printf("%v[%T]\n", s4, s4)
	fmt.Printf("%v[%T]\n", string(s4[7]), s4[7])

	//ostatni typ rune - reprezentuje UTF32 więc obsługuje znaki 8, 16 i 32 bajtowe więc trickygo tutorial - aliasy na typ int32
	r := 'a'
	var r2 rune = 'a'
	fmt.Printf("%v[%T]\n", r, r)
	fmt.Printf("%v[%T]\n", r2, r2)
}
