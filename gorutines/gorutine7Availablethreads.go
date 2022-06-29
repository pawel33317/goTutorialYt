package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(100) //ustawia jaka ma być ilość wątków
	//można zrobić więc 10 wątków na jednym fizycznym
	fmt.Printf("Available threads: %v", runtime.GOMAXPROCS(-1))
}
