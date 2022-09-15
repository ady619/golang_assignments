package main

import (
    "fmt"
)

type Account struct{
	UserId int
	Password string
	Balance int
}

var accountList []Account

func NavigateToLoginScreen(){
	fmt.Println("Please enter your user id: ")
	var UserId int
	fmt.Scan(&UserId)
	fmt.Println("Please enter your password: ")
	var Password string
	fmt.Scan(&Password)
	Login(UserId, Password)
}

func Login(id int, pwd string){
	found := false
	var User Account
	for _,item := range accountList{
		if item.UserId == id {
			found = true
			User = item
		}
	}

	if found {

		fmt.Println("Looged in user ",id,"\nPlease select an option from the menu below:")
		fmt.Println("w -> Withdraw money\nd -> Diposit Money\nr -> Request Balance\nq -> Quit\n")

		var option string
		fmt.Scan(&option)

		switch option {
		case "w":
			fmt.Println("Amount to withdraw: ")
			var wamount int
			fmt.Scan(&wamount)
			User.WithdrawMoney(wamount)
		case "d":
			fmt.Println("Amount to diposit: ")
			var damount int
			fmt.Scan(&damount)
			User.DipositMoney(damount)
		case "r":
			User.RequestBalance()
		default:
			QuitProgram()
		}
	} else {
		fmt.Println("Invalid userid or password")
		home()
	}
}

func NavigateToCreateAccountScreen(){
	fmt.Println("Please enter your New user id: ")
	var UserId int
	fmt.Scan(&UserId)
	fmt.Println("Please enter your New password: ")
	var Password string
	fmt.Scan(&Password)
	user := Account{UserId, Password, 0}
	user.CreateAccount()
}

func (user Account) CreateAccount(){
	accountList = append(accountList, user)
	fmt.Println(accountList)
	home()
}

func (ac Account) WithdrawMoney(amount int){
	ac.Balance = ac.Balance - amount
	fmt.Println("Money withdrawn: ",amount)
	home()
}

func (ac Account) DipositMoney(amount int){
	ac.Balance = ac.Balance + amount
	fmt.Println("Money Diposited: ",amount)
	home()
}

func (ac Account) RequestBalance(){
	fmt.Println("Balance: ",ac.Balance)
	home()
}

func QuitProgram(){
	home()
}

func main() {
	accountList = make([]Account, 0, 1)
	home()
}

func home(){
	fmt.Println("Hi! Welcome to Mr. Rohit ATM Machine!\nPlease select an option from the menu below:")
	fmt.Println("l -> Login\nc -> Create New Account\nq -> Quit\n")
	var option string
	fmt.Scan(&option)

	switch option {
	case "l":
		NavigateToLoginScreen()
	case "c":
		NavigateToCreateAccountScreen()
	case "q":
		QuitProgram()
	default:
		QuitProgram()
	}
}