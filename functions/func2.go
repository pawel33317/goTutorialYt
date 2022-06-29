package main

import "fmt"

func main() {
	fmt.Println("------------------------------by value")
	greetings1 := "Hello"
	name1 := "Pawel"
	sayGreeting___byValue(greetings1, name1)
	fmt.Println(name1)

	fmt.Println("------------------------------by pointer")
	greetings2 := "Hello"
	name2 := "Pawel"
	sayGreeting___byPointer(&greetings2, &name2)
	fmt.Println(name2)
	//mapy i slice podspodem mają pointery więc sa przekazywane przez pointer

	fmt.Println("------------------------------ variaduc parameters")
	sum("The sum is", 1, 2, 3, 4, 5, 6, 7)
}

func sayGreeting___byValue(greeting, name string) {
	fmt.Println(greeting, name)
	name = "Ted"
	fmt.Println(name)
}

func sayGreeting___byPointer(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}

func sum(text string, values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(text, result)
}
