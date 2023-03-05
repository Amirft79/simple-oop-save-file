package main

import (
	"fmt"
	"os"
)

type customer struct {
	customerName     string
	customerSureName string
}

func newCustomer() customer {
	c := customer{}
	return c
}

func (customer *customer) formatCustomer() string {
	fs := "the customerName is :" + customer.getCustomerName() + "and the customer sure name is " + customer.getCustomerSureName()
	return fs
}

func (customer *customer) setCustomerName(customerName string) {
	customer.customerName = customerName
}

func (customer *customer) setCustomerSureName(customerSureName string) {
	customer.customerSureName = customerSureName
}
func (customer *customer) getCustomerName() string {
	return customer.customerName
}
func (customer *customer) getCustomerSureName() string {
	return customer.customerSureName
}

func (customer *customer) saveFile() {
	data := []byte(customer.formatCustomer())
	err := os.WriteFile("userInput"+customer.getCustomerName()+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Print("file is saved")
}
