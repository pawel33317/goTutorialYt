package main

import "packageTestModule/commonPackage"

//go run -race src/packages/main.go src/packages/helper.go
//go run -race src/packages/*
func main() {
	printDate()
	commonPackage.PrintThreads()
}
