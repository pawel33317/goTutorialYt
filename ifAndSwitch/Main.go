package main

import (
	"fmt"
	"math"
)

func main() {
	if true {
		fmt.Println("Inside if")
	} //klamry obowiązkowe

	someMap := map[string]int{
		"Pawel":  15,
		"Marcin": 14,
		"Robert": 22,
	}

	if pop, ok := someMap["Pawel"]; ok { //pop, ok := someMap["Pawel"]; - to initializer
		fmt.Println(pop)
	}

	fmt.Println("----------------------------1")
	number := 50
	guess := 30

	if guess < 1 || returnTrue() { //wykona się bo guess > 1
		//short circuting ewaluuje ile musi
		fmt.Println("The guess must be between 1 and 100")
	} else if guess > 100 {
		fmt.Println("The guess must lower than ")
	} else {
		if guess < number {
			fmt.Println("Too low")
		}
		if guess > number {
			fmt.Println("Too high")
		}
		if guess == number {
			fmt.Println("Equal!!!")
		}
	}

	fmt.Println("----------------------------2 float comparsion")
	myNum := 0.123
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.00000001 {
		fmt.Println("Są mniej więcej takie same")
	}

	fmt.Println("----------------------------3 switch")

	switch 1 + 4 {
	case 1, 5, 7:
		fmt.Println("one or five or seven")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("not one and not two")
	} //break jest zawsze automatyczny

	ii := 10
	switch {
	case ii <= 10:
		fmt.Println("10 or less")
	case ii <= 20: //nie wykona sie bo pierwszy go zgarnie jak if i if else
		fmt.Println("20 or less")
	default:
		fmt.Println("other")
	}

	fmt.Println("----------------------------4")
	switch {
	case ii <= 10:
		fmt.Println("10 or less")
		fallthrough
	case ii > 20: //wykona się mimo że nie spełnia warunku przez fallthrough
		fmt.Println("20 or less") //chujnia
	default:
		fmt.Println("other")
	}

	fmt.Println("----------------------------TYPE SWITCH")

	var qq interface{} = 1.
	switch qq.(type) {
	case int:
		fmt.Println("int")
	case float64:
		fmt.Println("float64")
		break
		fmt.Println("float64 xxx") //nie wyświetli się
	case string:
		fmt.Println("string")
	default:
		fmt.Println("other type")
	}
}

func returnTrue() bool {
	fmt.Println("returning true")
	return true
}
