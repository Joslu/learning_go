package main

import "fmt"

func main() {

	var age int
	var adultYears int

	var agePointer *int

	age = 34 // Regular variable
	agePointer = &age

	fmt.Println("Actual age: ", *agePointer) //	Acceding to value stored on agePointer
	//fmt.Println("Adult years: ", getAdultYear(age))

	adultYears = getAdultYear(agePointer)
	fmt.Println("Adult years: ", adultYears)
}

func getAdultYear(age *int) int {

	return *age - 18
}
