package main

import "fmt"

func main() {

	// strings
	var nameOne string = "alph0n5e"
	var nameTwo = "me"
	var nameThree string
	fmt.Println(nameOne, nameTwo, nameThree)

	nameThree = "you"
	nameTwo = "us"
	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "them"
	fmt.Println(nameFour)

	// integers
	var ageOne int = 30
	var ageTwo = 31
	ageThree := 32
	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint8 = 255
	fmt.Println(numOne, numTwo, numThree)

	// floats
	var scoreOne float32 = -1.125
	var scoreTwo float64 = 987654332123456.54

	fmt.Println(scoreOne, scoreTwo)

}
