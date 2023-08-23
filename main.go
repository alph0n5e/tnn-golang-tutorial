package main

import "fmt"

func updateName(x *string) {
	*x = "wedge"
}

func main() {
	// Pointers
	name := "alph0n5e"
	m := &name

	fmt.Println("Memory address of 'name': ", m)
	fmt.Println("Value at memory address: ", *m)

	fmt.Println(name)
	updateName(m)
	fmt.Println(name)
}
