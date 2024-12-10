package main

import (
	"fmt"
	"time"
)

var balance float64 = 100

func main() {
	action := ""

	var amount float64
	receiversName := ""

	isRunning := true
	for isRunning {
		fmt.Print("\033[H\033[2J")
		fmt.Println("welcome to my bank choose the following actions " +
			"\npress: \n1 to send\n2 to withdraw\n3 to deposit\n4check balance\nq to quit")
		fmt.Print("enter action: ")

		fmt.Scan(&action)
		switch action {
		case "1":
			fmt.Print("\033[H\033[2J")
			fmt.Print("Enter the amount to send: ")
			fmt.Scan(&amount)
			fmt.Println("Enter the recipient name: ")
			fmt.Scan(&receiversName)

			sendMoney(amount, receiversName, balance)
		case "q":
			isRunning = false
			fmt.Println("quiting application!")
			fmt.Print("\033[H\033[2J")
		case "2":
			fmt.Print("\033[H\033[2J")
			withdrawFrom := ""
			agentNumber := 0
			fmt.Print("Withdraw from agent or atm \n1.agent\n2.atm: ")
			fmt.Scan(&withdrawFrom)
			if withdrawFrom == "agent" {
				fmt.Print("Enter the agent number to withdraw from: ")
				fmt.Scan(&agentNumber)
				fmt.Print("Enter the amount to withdraw: ")
				fmt.Scan(&amount)
				withdrawFunds(amount, agentNumber, balance)
			} else if withdrawFrom == "atm" {
				fmt.Print("Enter the Atm number to withdraw from: ")
				fmt.Scan(&agentNumber)
				fmt.Print("Enter the amount to withdraw: ")
				fmt.Scan(&amount)
				withdrawFunds(amount, agentNumber, balance)
			}
		case "3":
			fmt.Print("Enter the amount to deposit: ")
			fmt.Scan(&amount)
			depositFunds(amount, balance)
		case "4":
			fmt.Println("your balance is ", balance)
		}

	}

}
func sendMoney(amount float64, receiverName string, balance float64) {
	type send struct {
		amount   float64
		receiver string
	}
	if amount < balance && balance > 0 {
		balance -= amount
		transaction := send{amount: amount, receiver: receiverName}
		fmt.Println("Confirmed ksh ",

			transaction.amount, "sent to ",
			transaction.receiver,
			"at ", time.Now(),
			"your balance is Kshs ", balance)
	} else {
		fmt.Println("cannot send more than your account balance! \nplease top up or send a lower amount because your balance is", balance)
	}

}

func withdrawFunds(amount float64, agentNumber int, balance float64) {
	type withdraw struct {
		amount      float64
		agentNumber int
	}
	if amount < balance && balance > 0 {
		balance -= amount
		transaction := withdraw{amount: amount, agentNumber: agentNumber}
		fmt.Println("Confirmed ksh ",
			transaction.amount, "withdrawn at ",
			transaction.agentNumber,
			"at ", time.Now(),
			"your balance is Kshs ", balance)
	} else {
		fmt.Println("upuzii")

	}

}
func depositFunds(amount float64, balance float64) {
	if amount > 0 {
		balance += amount
		fmt.Println("Confirmed ksh ",
			amount, "deposited to your account at ",
			time.Now(),
			"your balance is Kshs ", balance)
	} else {
		fmt.Println("cannot deposit funds less than 0")
	}

}
