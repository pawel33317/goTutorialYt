package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := 42
	var b *int = &a //b:=&a
	a = 27
	*b = 333
	fmt.Println(a, *b)
	fmt.Println(&a, b)

	arr := []int32{1, 2, 3} //normalnie int u mnie ma 64
	arb := &arr[0]
	arc := &arr[1] // - 4 -> error, go nie pozwala na arytmetykę
	fmt.Printf("%v %p %p \n", arr, arb, arc)

	unsafePtr := unsafe.Pointer(&arr[1])
	itemSize := unsafe.Sizeof(arr[1])
	fmt.Printf("%v, %v\n", unsafePtr, itemSize)
	// res := *(*int32)(unsafe.Add(unsafePtr, itemSize)) //od 1.17
	// fmt.Printf("%v\n", res)

	fmt.Println("-----------------")

	var ms1 myStruct
	ms1 = myStruct{foo: 42}
	fmt.Println(ms1)

	var ms2 *myStruct
	ms2 = &myStruct{foo: 42}
	fmt.Println(ms2)

	fmt.Println("-----------------ms3")
	var ms3 *myStruct
	fmt.Println(ms3)    //nil
	ms3 = new(myStruct) // nie da się użyć składni inicjalizyjnej
	fmt.Println(ms3)
	(*ms3).foo = 42
	fmt.Println((*ms3).foo)
	fmt.Println(ms3)
	ms3.foo = 555 //!!! kompilator to umie interpretować
	fmt.Println(ms3)

	fmt.Println("-----------------")
	x := [...]int{1, 2, 3}
	y := x //to jest kopia array więc zmieni tylko w jednej
	fmt.Println(x, y)
	x[1] = 42
	fmt.Println(x, y)

	x2 := []int{1, 2, 3}
	y2 := x //to vector (slice) a nie array więc zmieni w obu bo to pointer
	fmt.Println(x2, y2)
	x2[1] = 42
	fmt.Println(x2, y2)

	//sherowanie slice jest przez pointery

	fmt.Println("-----------------mapa")

	mapa := map[string]string{"foo": "bar", "baz": "buz"}
	mapb := mapa
	fmt.Println(mapa, mapb)
	mapa["foo"] = "qux"
	fmt.Println(mapa, mapb)
	//mapy oczywiście też przez pointery
}
