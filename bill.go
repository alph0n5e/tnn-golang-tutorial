package main

import (
	"fmt"
	"os"
	"strings"
)

type bill struct {
	name  string
	items map[string]float32
	tip   float32
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float32{},
		tip:   0,
	}
	return b
}

// Format the bill (receiver function)
func (b bill) format() string {
	fs := "Bill breakdown:\n"
	var total float32 = 0

	// List items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-12v ...$%v\n", k, v)
		total += v
	}

	// Tip
	fs += "+\n"
	fs += fmt.Sprintf("%-12v ...$%0.2f\n", "tip:", b.tip)
	total += b.tip

	// Total
	fs += "---\n"
	fs += fmt.Sprintf("%-12v ...$%0.2f", "total:", total)

	return fs
}

// Update tip
func (b *bill) updateTip(tip float32) {
	b.tip = tip
}

// Add item
func (b *bill) addItem(name string, price float32) {
	b.items[name] = price
}

// Save bill
func (b *bill) save() {
	fileName := strings.ReplaceAll(b.name, " ", "-")
	data := []byte(b.format())
	err := os.WriteFile("bills/"+fileName+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill was saved to file")
}
