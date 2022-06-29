package main

import "fmt"

func main() {
	myInt := IncCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

type Incrementer interface {
	Increment() int
}

type IncCounter int

func (ic *IncCounter) Increment() int {
	*ic++
	return int(*ic)
}
