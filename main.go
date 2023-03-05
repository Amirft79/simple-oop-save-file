package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readValue(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func FillCustomer(customer *customer) customer {
	reader := bufio.NewReader(os.Stdin)
	option, _ := readValue("which item do you want add a(cusotmerNmae) b(customerSureName): ", reader)
	switch option {
	case "a":
		customerName, _ := readValue("enter your customerName : ", reader)
		customer.setCustomerName(customerName)
		FillCustomer(customer)
	case "b":
		customerSureName, _ := readValue("ENTER CUSTOMER sureName : ", reader)
		customer.setCustomerSureName(customerSureName)
		FillCustomer(customer)
	case "s":
		customer.saveFile()
	default:
		FillCustomer(customer)
	}

	return *customer

}

func main() {
	customer := newCustomer()
	customer = FillCustomer(&customer)
	fmt.Println(customer)
	fmt.Println(customer.getCustomerName())
	fmt.Println(customer.getCustomerSureName())

}
