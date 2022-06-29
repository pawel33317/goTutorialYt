package main

func main() {
	//używać wielu małych interfejsów
	//najlepiej jeden interfacej na jednąmetodę
	//io.Writer, io.Reader, interface{}

	//nie eksportować interfejsów do typów, które będą używane
	//np db obj, przez to user może sobie sam zdefiniować interfejs
	//i użyć go to testów i oląć metody, których nie używa
	//i tak ma pełny obiekt więc po co interfejs

	//eksportować interfejsy dla typów używanych przez package
	//których użytkownik nie będie widział

	//projektować funkcje i metody żeby używały interfejsów
	//kiedy tylko to możliwe zamiast konkretnych typów
}
