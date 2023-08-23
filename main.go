package main

import "fmt"

func updateName(x string) {
	x = "wedge"
}

func updateMenu(m map[string]float32) {
	m["tea"] = 2.5
}

func main() {
	// Non-pointer values
	name := "alph0n5e"
	updateName(name)
	fmt.Println(name)

	// Pointer wrapper values
	menu := map[string]float32{
		"coffee": 1.99,
		"pie":    4.99,
	}
	fmt.Println(menu)
	updateMenu(menu)
	fmt.Println(menu)
}
