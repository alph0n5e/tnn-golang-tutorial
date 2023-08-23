package main

import "fmt"

func main() {
	myBill := newBill("My Bill")

	myBill.addItem("coffee", 1.99)
	myBill.addItem("tea", 2.99)
	myBill.addItem("pie", 3.99)
	myBill.addItem("cake", 2.50)

	myBill.updateTip(1.2)

	fmt.Println(myBill.format())
}
