package main

import "fmt"

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))
}

type Writer interface { //opisują nie dane a zachowania
	Write([]byte) (int, error)
} //jak jest interfejs z jedną metodą to zwykle go nazywamy jak metoda + er

type ConsoleWriter struct{}

//implementacja interfejsu nie jest jawna
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
