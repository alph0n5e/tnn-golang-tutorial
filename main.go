package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func promptName() string {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name: ", reader)
	name = strings.TrimSpace(name)
	return name
}

func createBill() bill {
	name := promptName()

	if name == "" {
		fmt.Println("Bill name cannot be empty!")
		name = promptName()
	}

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price ($): ", reader)
		p, err := strconv.ParseFloat(price, 32)
		if err != nil {
			fmt.Println("The price must be a number!")
			promptOptions(b)
		}
		b.addItem(name, float32(p))
		fmt.Println("Item added: ", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 32)
		if err != nil {
			fmt.Println("The tip must be a number!")
			promptOptions(b)
		}
		b.updateTip(float32(t))
		fmt.Println("Tip added: ", tip)
		promptOptions(b)
	case "s":
		b.save()
	default:
		fmt.Printf("Unknown option \"%v\"\n", opt)
		promptOptions(b)
	}
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
}
