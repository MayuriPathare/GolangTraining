package main

import (
	"fmt"
)

type userFunctionality interface{
	deposit(amount float64) 
	withdraw(amount float64) 
	checkBalance() float64
}

type currentAccount struct{
	balance float64
}

type savingAccount struct{
	balance float64
	interestRate float64

}

func (ca *currentAccount) deposit(amount float64){
	fmt.Printf("Amount Deposited: Rs. %.2f", amount)
	ca.balance += amount
	fmt.Printf("Available balance:Rs. %.2f\n", ca.balance)
}
func (ca *currentAccount) withdraw(amount float64) {
	fmt.Printf("Amount withdraw %f\n", amount)
 	ca.balance-=amount
 	fmt.Printf("Available balance:Rs. %.2f\n", ca.balance)

}
func (ca *currentAccount) checkBalance() float64 {
	fmt.Printf("Balance of Current  account is: Rs. %.2f\n", ca.balance)
	return ca.balance
}

func (sa *savingAccount) deposit(amount float64) {
	sa.balance+=(amount+sa.interestRate*amount)
	fmt.Printf("Amount Deposited: Rs. %.2f, Available balance:Rs. %.2f\n", amount, sa.balance)
}
func (sa *savingAccount) withdraw(amount float64) {
    sa.balance-=amount
    fmt.Printf("Amount Withdrawn: Rs. %.2f, Avalable balance: Rs. %.2f\n", amount, sa.balance)
}
func (sa *savingAccount) checkBalance() float64 {
	fmt.Printf("Balance of Saving account is: Rs. %.2f\n", sa.balance)
    return sa.balance
}


func main(){

	var Acc1 userFunctionality = &currentAccount{balance:1000}
	var Acc2 userFunctionality = &savingAccount{balance:1000, interestRate:10}
    fmt.Println("Balance in current account")
	Acc1.checkBalance()
	Acc1.deposit(1000)
	Acc1.checkBalance()
	Acc2.withdraw(1000)
}
