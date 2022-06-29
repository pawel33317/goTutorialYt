package main

import "fmt"

func main() {
	grades := [3]int{56, 67, 78}    //w ciągłej przestrzeni
	grades2 := [...]int{56, 67, 78} //w ciągłej przestrzeni
	fmt.Println("Grades: ", grades)
	fmt.Println("Grades: ", grades2)

	var students [3]string //może być każdego typu ale cała tego samego
	fmt.Println("Students: ", students)
	students[0] = "Pawel"
	fmt.Println("Students: ", students)
	fmt.Println("Students len: ", len(students))

	var idMatrix [3][3]int = [3][3]int{{0, 0, 0}, {1, 1, 1}, {2, 2, 2}}
	idMatrix[0] = [...]int{5, 5, 5}
	idMatrix[1] = [3]int{9, 9, 9}
	fmt.Println("idMatrix: ", idMatrix)

	q1 := [...]int{1, 2, 3, 4}
	q2 := q1  //!!!! ROBI KOPIE
	q3 := &q2 //pointer - NIE ROBI KOPII
	q2[0] = 888
	fmt.Println(q1)
	fmt.Printf("Array type is %T\n", q1)
	fmt.Printf("ArrayPointer type is %T\n", &q1)
	fmt.Println(q2)
	fmt.Println(*q3)
	fmt.Println(q3)
	fmt.Println(&q3)

	fmt.Println("-----------SLICES------------")

	q1Slice := []int{1, 2, 3}
	fmt.Println("Slice val: ", q1Slice)
	fmt.Printf("Slice type is %T\n", q1Slice)
	fmt.Printf("Slice Pointer type is %T\n", &q1Slice)
	fmt.Printf("Slice len is %v\n", len(q1Slice))
	fmt.Printf("Slice capacity is %v\n", cap(q1Slice))
	q2Slice := q1Slice
	q1Slice[0] = 999 //!!!!!!!!! slice już nie są kopiowane i działają jak referencje
	fmt.Println("Slice1 val: ", q1Slice)
	fmt.Println("Slice2 val: ", q2Slice)
	//arrays muszą mieć znany rozmiar w compile time więc są slices

	w1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	w2 := w1[:]   //kopiuje całą
	w3 := w1[3:]  //od 4-tygo elementu ([3] indeksu)
	w4 := w1[:6]  //pierwsze 6 elementów (bez [6])
	w5 := w1[3:6] //4,5,6 ([3][4][5])
	//[inclusive, exclusive]
	fmt.Println(w1)
	fmt.Println(w2)
	fmt.Println(w3)
	fmt.Println(w4)
	fmt.Println(w5)

	fmt.Println("-----------SLICES using make------------")
	t1Slice := make([]int, 3, 54) // opcjonalne capacity
	fmt.Println("Slice val: ", t1Slice)
	fmt.Printf("Slice type is %T\n", t1Slice)
	fmt.Printf("Slice len is %v\n", len(t1Slice))
	fmt.Printf("Slice capacity is %v\n", cap(t1Slice))

	fmt.Println("-----------SLICES...------------")
	y1Slice := []int{} // opcjonalne capacity
	fmt.Println("Slice val: ", y1Slice)
	fmt.Printf("Slice len is %v\n", len(y1Slice))
	fmt.Printf("Slice capacity is %v\n", cap(y1Slice))

	y1Slice = append(y1Slice, 33)
	fmt.Println("Slice val: ", y1Slice)
	fmt.Printf("Slice len is %v\n", len(y1Slice))
	fmt.Printf("Slice capacity is %v\n", cap(y1Slice))

	y1Slice = append(y1Slice, 33, 44, 55, 66, 77, 11) //funkcja wariadyczna
	fmt.Println("Slice val: ", y1Slice)
	fmt.Printf("Slice len is %v\n", len(y1Slice))
	fmt.Printf("Slice capacity is %v\n", cap(y1Slice))

	y1Slice = append(y1Slice, []int{1, 2, 3}...) //coś jak fold expresion
	fmt.Println("Slice val: ", y1Slice)
	fmt.Printf("Slice len is %v\n", len(y1Slice))
	fmt.Printf("Slice capacity is %v\n", cap(y1Slice))

	fmt.Println("-----------SLICES stack operations------------")
	sliceStack := []int{1, 2, 3, 4, 5, 6, 7}
	sliceStack = sliceStack[:len(sliceStack)-1] //usuniecie najnowszego elementu
	fmt.Println("Slice val: ", sliceStack)

	sliceStack2 := append(sliceStack[:2], sliceStack[3:]...)
	//!!!modyfikuje też referencje do tego czyli sliceStack

	fmt.Println("Slice val: ", sliceStack2)
	fmt.Println("Slice val: ", sliceStack) //!!!undefined behavior
}
