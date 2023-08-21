package main

import "fmt"

var points = []int{10, 12, 15, 20}

func greet(name string) {
	fmt.Printf("Greetings %v!\n", name)
}

func showScore() {
	fmt.Println("You scored this many points: ", score)
}
