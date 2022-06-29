package main

import "fmt"

func main() {
	fmt.Println("------------------------------ return val")
	fmt.Println("The sum is", sum(1, 2, 3, 4, 5, 6, 7))

	fmt.Println("------------------------------ return LOCAL pointer")
	fmt.Println("The sum is", *sum2(1, 2, 3, 4, 5, 6, 7))

	fmt.Println("------------------------------ named return value")
	fmt.Println("The sum is", sum3(1, 2, 3, 4, 5, 6, 7))

	fmt.Println("------------------------------ return multiple values")

	res, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func sum(values ...int) int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

//UMIE zwrócić adres do lokalnej zmiennej
//automatycznie to wychwytuje i przerzuca na HEAP
func sum2(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

func sum3(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
