package main

import (
	"fmt"
	"time"
)

func main() {

	//grean thread
	//standardowe języki używają często west threads - operating system
	//threads które często są duże np 1MB ramu i długo trwa ich tworzenie
	//mają indywidualny call stack - te z innych języków
	//to wzięli z erlanga - green threads
	//abstrakcja wątku - gorutyna
	//w środku go jest scheduler, który mapuje gorutiny na wątki,
	//na określony czas na określone wątki procesora
	//go rutyny są tanie w tworzeniu i często mogą być ich tysiące
	//w innych językach to raczej niemożliwe
	go sayHello()

	//nic się nie wyświetla bo główna się skończyła
	time.Sleep(100 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello")
}
