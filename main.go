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

func FillAccount(account *account) account {
	reader := bufio.NewReader(os.Stdin)
	output, _ := readValue("please selcet your account options a(create account) b(save account) : ", reader)
	switch output {
	case "a":
		fmt.Print("account number :")
		read, _ := readValue("input your account number", reader)
		account.setAccountNumber(read)
		readedate, _ := readValue("input your account date", reader)
		account.setAccountOpenDate(readedate)
		readeId, _ := readValue("input  your account customer id ", reader)
		account.setAccountCustomerId(readeId)
		FillAccount(account)
	case "s":
		account.saveFile()
	}
	return *account

}

func main() {
	customer := newCustomer()
	customer = FillCustomer(&customer)
	fmt.Println(customer)
	fmt.Println(customer.getCustomerName())
	fmt.Println(customer.getCustomerSureName())
	account := newAccount()
	account = FillAccount(&account)
	fmt.Println(account)
	fmt.Println(account.getAccountNumber())
}
