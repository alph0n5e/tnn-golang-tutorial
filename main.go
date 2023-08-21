package main

import (
	"fmt"
	"math"
	"strings"
)

func sayGreeting(name string) {
	fmt.Printf("Good morning %v!\n", name)
}

func sayBye(name string) {
	fmt.Printf("Goodbye %v!\n", name)
}

func cycleNames(names []string, f func(string)) {
	for _, name := range names {
		f(name)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func getInitials(names string) (string, string) {
	s := strings.ToUpper(names)
	ss := strings.Split(s, " ")

	var initials []string
	for _, v := range ss {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	name := "Antoine"
	sayGreeting(name)
	sayBye(name)
	cycleNames([]string{"Antoine", "Robert", "alph0n5e"}, sayGreeting)
	fmt.Println(circleArea(2))

	fN, lN := getInitials("Antoine ROBERT")
	fmt.Println(fN, lN)
}
