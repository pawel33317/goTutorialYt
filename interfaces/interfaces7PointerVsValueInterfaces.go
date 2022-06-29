package main

import "fmt"

func main() {
	//var wc WriterCloser = myWriterCloser{} - no compile
	//dodałem "*" na  func (mwc *myWriterCloser) Write...
	//i w związku że teraz przyjmuje pointer nie umie przekonwertować
	//Write method has pointer receiver

	//teraz interfejs nie ma podstawionej wartości
	//tylko wskazuje na obiekt
	//lepsza opcja implementacja interfejsu przez pointer
	//bo teraz metody mogą obsłużyć go i przez pointer i wartość
	//tak jak poniżej Write i Close
	var wc WriterCloser = &myWriterCloser{}
	fmt.Println(wc)
}

type Writer interface {
	Write([]byte) (int, error)
}
type Closer interface {
	Close() error
}
type WriterCloser interface {
	Writer
	Closer
}

type myWriterCloser struct{}

//obsługuje i value receivera i pointer receivera
func (mwc *myWriterCloser) Write(data []byte) (int, error) {
	return 0, nil
}

//obsługuje tylko value receivera
func (mwc myWriterCloser) Close() error {
	return nil
}
