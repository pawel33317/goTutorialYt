package main

import "fmt"

//w zależności czy z dużej czy z małej funkcja jest internal lub external
func main() {
	fmt.Println("Hello, go")

	sayMessage("Hello GO!")

	sayMessage("Hello GO!")
	for i := 0; i < 5; i++ {
		sayMessage2("Hello Gooo", i)
	}
}

func sayMessage(msg string) {
	fmt.Println(msg)
}

func sayMessage2(msg string, idx int) {
	fmt.Println(idx, msg)
}
