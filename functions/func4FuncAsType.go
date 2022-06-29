package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello Go !!!")
	}()

	for i := 0; i < 5; i++ {
		func() { //byłby problem przy wywołaniu asynchronicznym funkcji
			fmt.Println(i)
		}()
	}

	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}

	f := func() {
		fmt.Println("Gello from f")
	}
	var f2 func() = func() {
		fmt.Println("Gello from f")
	}
	f()
	f2()

	var divide func(float64, float64) (float64, error)
	divide = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Cannot divide by zero")
		}
		return a / b, nil
	}
	res, err := divide(5., 3.)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	fmt.Println("Function type is %T", divide)
}
