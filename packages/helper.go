package main

import (
	"fmt"
	"time"
)

func printDate() {
	fmt.Println("Current time: ", time.Now().Format("2006.01.02 15:04:05"))
}
