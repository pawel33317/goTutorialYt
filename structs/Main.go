package main

import (
	"fmt"
	"reflect"
)

//z dużej jest External, ale pola są prywatne dla extern użycia bo są z małej
type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Pawel",
		companions: []string{
			"aaa",
			"bbb",
			"ccc",
			"ddd",
		},
	}

	bDoctor := Doctor{
		3,
		"Pawel",
		[]string{
			"aaa",
		},
	}

	fmt.Printf("aDoctor value: %v, type %T\n", aDoctor, aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(bDoctor)

	anonymopusStructObj := struct{ name string }{name: "Pawel"}
	anotherAnonymousStructObj := anonymopusStructObj
	anotherAnonymousStructObj.name = "Marian"
	fmt.Println(anonymopusStructObj)
	fmt.Println(anotherAnonymousStructObj)

	fmt.Println("---------------------------")
	type Animal struct {
		Name   string
		Origin string
	}
	type Bird struct {
		Animal   //nie jako pole tylko sama nazwa - oznacza że ma wbudowane Animal embe
		SpeedKPH float32
		CanFly   bool
	}

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)
	fmt.Println(b.Name)
	//w GO nie ma dziedziczenia a jest kompozycja
	//GO automatycznie forwarduje takie requesty do typu embeded
	//to kompozycja więc Bird nie jest Animal -> później bedzie o interfejsach

	b2 := Bird{
		Animal:   Animal{Name: "Skowronek", Origin: "Polska"},
		SpeedKPH: 66,
		CanFly:   true,
	}
	fmt.Println(b2)

	fmt.Println("--------------tagi -------------")
	type Animal2 struct {
		Name   string `required max: "100` //obowiązkowa i ustawiony max size
		Origin string
	}
	t := reflect.TypeOf(Animal2{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
