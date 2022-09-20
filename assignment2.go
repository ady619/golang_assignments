package main

import (
    "fmt"
)

// Given two integers a and b, return the sum of the two integers 
// without using the operators + and -.

func main(){
	l := SumWithoutOPperators(5, 7)
	fmt.Println(l)
}

func SumWithoutOPperators(a int,b int)int{
	m := make([]int, a)
	n := make([]int, b)
	m = append(m, n...)
	return len(m)
}