package main

import "fmt"

func main() {
	population := map[string]int{
		"Poland":  37,
		"Germany": 78,
		"England": 66,
	}

	fmt.Printf("Population val: %v, type: %T\n", population, population)
	fmt.Println("Poland population: ", population["Poland"])
	population["France"] = 58
	population["Poland"] = 36
	//kluczem musi być coś testowalne do equality

	newMap := make(map[string]int)
	newMap = population
	fmt.Println(newMap)

	delete(population, "Poland")
	fmt.Println(newMap)

	fmt.Println("Poland population: ", population["Poland"]) //0 bo brak w liście

	polandPopulation, ok := population["Poland"]
	fmt.Println(ok, " - ", polandPopulation)

	francePopulation, ok := population["France"]
	fmt.Println(ok, " - ", francePopulation)

	fmt.Println(len(population))

	//kolejność w mapie nie jest gwarantowana, a w slice jest
}
