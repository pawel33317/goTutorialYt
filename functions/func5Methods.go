package main

import "fmt"

func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
	fmt.Println("The new name is: ", g.name)
	g.greetPtr()
	fmt.Println("The new name is: ", g.name)
}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() { //value receiver robi kopię
	fmt.Println(g.greeting, g.name)
	g.name = ""
}

func (g *greeter) greetPtr() { //pointer receiver robi kopię
	fmt.Println(g.greeting, g.name)
	g.name = ""
}
