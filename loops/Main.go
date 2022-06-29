package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//w go nie ma operatora ,
	//i++ to nie wyrażenie jak w innych językach a statement
	for i, j := 0, 0; i < 5; i, j = i+2, j+5 {
		fmt.Println(i, " - ", j)
	}

	fmt.Println("-----------------------------")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		i++
	}

	fmt.Println("-----------------------------")
	k := 0
	for ; k < 5; k++ {
		fmt.Println(k)
	}

	fmt.Println("-----------------------------")
	k2 := 0
	for k2 < 5 {
		fmt.Println(k2)
		break
	}

	fmt.Println("----------------------------- infinite")
	for {
		fmt.Println(999)
		break
	}

	fmt.Println("----------------------------- continue")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("----------------------------- goto")
LabelLoop: //musi być przed jakąś pętlą przet printem się nie dało
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i*j > 10 {
				break LabelLoop
			}
		}
	}

	fmt.Println("----------------------------- kolekcje i for")
	s := [...]int{1, 2, 3, 4, 5}
	s2 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Tab %T, slice %T\n", s, s2)

	fmt.Println(s2)

	for k, v := range s2 {
		fmt.Println(k, v)
	}

	someMap := map[string]int{
		"Pawel":  14,
		"robert": 12,
	}

	for k, v := range someMap {
		fmt.Println(k, v)
	}

	for k := range someMap {
		fmt.Println(k)
	}

	for _, v := range someMap {
		fmt.Println(v)
	}
}
