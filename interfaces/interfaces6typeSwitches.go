package main

import (
	"fmt"
)

func main() {
	var i interface{} = true //dla 0 jest int
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is an string")
	default:
		fmt.Println("i is unknown type")
	}
}
