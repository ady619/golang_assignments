package main

import (
    "fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"log"
)

type Account struct{
	UserId int `json:"UserId"`
	Password string
	Balance int
}

var accountList *[]Account
var dataFile string

func EnterCredentials(){
	fmt.Println("Please enter your user id: ")
	var UserId int
	fmt.Scan(&UserId)
	fmt.Println("Please enter your password: ")
	var Password string
	fmt.Scan(&Password)
	Login(UserId, Password)
}

func Login(id int, pwd string){
	// found := false
	// var user Account
	index := -1
	for i,item := range *accountList{
		if item.UserId == id {
			index = i
		}
	}
	user := &(*accountList)[index]
	// fmt.Println(*accountList)

	if index > -1 {
		fmt.Println("\nLooged in user :",(*user).UserId)
		UserHomeScreen(user)
	} else {
		fmt.Println("\n!! Invalid userid or password !!\n")
		home()
	}
}

func UserHomeScreen(user *Account){
	
	START:
	fmt.Println("\nPlease select an option from the menu below:")
	fmt.Println("w -> Withdraw money\nd -> Diposit Money\nr -> Request Balance\nl -> Log out\n")

	var option string
	fmt.Scan(&option)

	switch option {
	case "w":
		fmt.Println("Amount to withdraw: ")
		var amount int
		fmt.Scan(&amount)
		(*user).WithdrawMoney(amount)
	case "d":
		fmt.Println("Amount to diposit: ")
		var amount int
		fmt.Scan(&amount)
		(*user).DipositMoney(amount)
	case "r":
		(*user).RequestBalance()
	case "l":
		home()
	default:
		fmt.Println("\n!! Invalid Option selected !!\n")
		goto START
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
	CreateAccount(&user)
}

func CreateAccount(user *Account){
	*accountList = append(*accountList, *user)
	fmt.Println("\nLooged in user :",(*user).UserId)
	content,err := json.Marshal(accountList)
	if err != nil{
		fmt.Println(err)
	}
	
	if err = ioutil.WriteFile(dataFile, content, 0644); err != nil{
		log.Fatal(err)
	}
	fmt.Println(accountList)
	UserHomeScreen(user)
}

func (user *Account) WithdrawMoney(amount int){
	(*user).Balance = (*user).Balance - amount
	fmt.Println("Money withdrawn: ",amount)
	fmt.Println(accountList)
	UserHomeScreen(user)
}

func (user *Account) DipositMoney(amount int){
	(*user).Balance = (*user).Balance + amount
	fmt.Println("Money Diposited: ",amount)
	fmt.Println(accountList)
	UserHomeScreen(user)
}

func (user *Account) RequestBalance(){
	fmt.Println("Balance: ",(*user).Balance)
	fmt.Println(accountList)
	UserHomeScreen(user)
}

func QuitProgram(){
	content,err := json.Marshal(accountList)
	fmt.Println(accountList)
	fmt.Println(string(content))
	if err != nil{
		log.Fatal(err)
	}
	
	if err = ioutil.WriteFile(dataFile, content, 0644); err != nil{
		log.Fatal(err)
	}
	os.Exit(3)
}

func main() {
	dataFile = "dataFile.json"
	file,err := os.Open(dataFile)
	if err != nil{
		_,err := os.Create(dataFile)
		if err != nil{
			panic(err)
		}
		main()
	}
	defer file.Close()

	var userList []Account
	if err = json.NewDecoder(file).Decode(&userList);err != nil {
		err = ioutil.WriteFile(dataFile, []byte("[]"), 0644)
		if err != nil{
			log.Fatal(err)
		}
		main()
	}

	list := make([]Account, 0, 1)
	accountList = &list
	*accountList = append(list, userList...)

	fmt.Println(*accountList)
	home()
}

func home(){
	fmt.Println("\nHi! Welcome to Aditya ATM Machine!")
	START:
	fmt.Println("Please select an option from the menu below:")
	fmt.Println("l -> Login\nc -> Create New Account\nq -> Quit")
	var option string
	fmt.Scan(&option)

	switch option {
	case "l":
		EnterCredentials()
	case "c":
		NavigateToCreateAccountScreen()
	case "q":
		QuitProgram()
	default:
		fmt.Println("\n!! Invalid Option selected !!\n")
		goto START
	}
}