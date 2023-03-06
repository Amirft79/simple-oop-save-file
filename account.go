package main

import (
	"fmt"
	"os"
)

type account struct {
	accountNumber     string
	accountOpenDate   string
	accountCustomerId string
}

func newAccount() account {
	account := account{}
	return account
}

func (account *account) formatAccount() string {
	fs := "the accountNumber is :" + account.getAccountNumber() + "and the account OpenDate  is " + account.getAccountOpenDate() + "" +
		"and the account customer id is " + account.getAccountCustomerId()
	return fs
}

func (account *account) setAccountNumber(accountNumber string) {
	account.accountNumber = accountNumber
}
func (account *account) setAccountOpenDate(accountOpenDate string) {
	account.accountOpenDate = accountOpenDate
}
func (account *account) setAccountCustomerId(accountCustomerId string) {
	account.accountCustomerId = accountCustomerId
}

func (account *account) getAccountNumber() string {
	return account.accountNumber
}
func (account *account) getAccountOpenDate() string {
	return account.accountOpenDate
}
func (account *account) getAccountCustomerId() string {
	return account.accountCustomerId
}

func (account *account) saveFile() {
	data := []byte(account.formatAccount())
	err := os.WriteFile("userInput"+account.getAccountNumber()+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Print("file is saved")
}
