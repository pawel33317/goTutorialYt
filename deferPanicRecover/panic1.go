package main

import "fmt"

//exception w go to panic

func main() {
	fmt.Println("start")
	defer fmt.Println("this was deffered") //to się wykona mimo rzucenia wyjątku
	panic("something bad happened")
	fmt.Println("end")

	/*a, b := 1, 0
	ans := a / b
	fmt.Println(ans)*/
}
