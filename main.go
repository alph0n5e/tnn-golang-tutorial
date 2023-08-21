package main

import "fmt"

func main() {

	// Arrays

	var ages [3]int = [3]int{20, 25, 30}
	names := [4]string{"Yoshi", "Mario", "Peach", "Bowser"}
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	names[1] = "Luigi"
	fmt.Println(names, len(names))

	// Slices

	var scores = []int{100, 50, 60}

	fmt.Println(scores, len(scores))

	scores[2] = 25
	fmt.Println(scores, len(scores))

	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	// Slice ranges

	rangeOne := names[1:3]
	rangeTwo := names[1:]
	rangeThree := names[:3]
	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

	rangeOne = append(rangeOne, "Waluigi") // Will append to 'names' as well
	fmt.Println(rangeOne)

	rangeTwo = append(rangeTwo, rangeThree...)
	fmt.Println(rangeTwo)
	fmt.Println(names)
}
