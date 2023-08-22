package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":  4.99,
		"pie":   7.99,
		"salad": 6.99,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])
	for k, v := range menu {
		fmt.Printf("menu[\"%v\"] = %v\n", k, v)
	}

	phonebook := map[int]string{
		234567: "mario",
		123456: "luigi",
		345678: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[123456])

	phonebook[123456] = "bowser"
	fmt.Println(phonebook)
}
