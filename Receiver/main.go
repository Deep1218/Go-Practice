package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// func main() {
// 	mybill := newBill("mario's bill")

// 	mybill.addItem("onion soup", 400.50)
// 	mybill.addItem("veg pie", 80.95)
// 	mybill.addItem("toffee pudding", 40.95)
// 	mybill.addItem("coffee", 300.25)

// 	mybill.updateTip(1000)

//		fmt.Println(mybill.format())
//	}
func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, error := r.ReadString('\n')

	return strings.TrimSpace(input), error
}
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name: ", reader)
	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)
	return b
}
func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option from (a -add item, s - save bill, t - add tip): ", reader)
	// fmt.Println(opt)
	switch opt {
	case "a":
		itemName, _ := getInput("Enter Item Name: ", reader)
		itemPrice, _ := getInput("Enter Item Price: ", reader)
		p, err := strconv.ParseFloat(itemPrice, 64)
		if err == nil {
			b.addItem(itemName, p)
			fmt.Println("Item added - ", itemName, itemPrice)
		} else {
			fmt.Println("Price should be number only.")
		}
		promptOptions(b)
	case "s":
		b.savebill()
		fmt.Println(b.name, " - Bill has been saved")
	case "t":
		tip, _ := getInput("Enter Tip Amount (₹): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err == nil {
			b.updateTip(t)
			fmt.Println("Tip - ", tip)
		} else {
			fmt.Println("Tip amount should be number only.")
		}
		promptOptions(b)
	default:
		fmt.Println(":( Invalid Input!! Please enter the valid choose")
		promptOptions(b)
	}
}
func main() {
	mybill := createBill()
	promptOptions(mybill)
}
