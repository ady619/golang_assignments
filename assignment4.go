package main

import (
    "fmt"
	"time"
	// "strings"
)

/*
Generate the Series 1A2B3C4D5E....... using  goroutine and channel  
.Routine 1 shall generate A,B C D and routine 2 shall generate 1 ,2,3,4,5,6....... 
Both this output join together in the output 
*/

func main(){
	ch := make(chan string, 0)
	go getNumbers(ch)
	go getAplhabets(ch)
	
	for n := range ch {
		fmt.Print(n)
	}
	var input string
	fmt.Scan(&input)
}

func getAplhabets(ch chan string){
	for c := 'A'; c <= 'Z'; c++ {
        ch<- fmt.Sprintf("%c", c)
		time.Sleep(100*time.Millisecond)
		// t:=<-ch
		// fmt.Printf(t)
		// fmt.Println("go getAplhabets")
    }
	return
}

func getNumbers(ch chan string){
	for i := 1; i <= 26; i++ {
        ch<- fmt.Sprintf("%v", i)
		time.Sleep(100*time.Millisecond)
		// t:=<-ch
		// fmt.Println(t)
		// fmt.Println("go getNumbers")
    }
	return
}