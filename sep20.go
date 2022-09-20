package main

import (
    "fmt"
)

func main(){
	l := SumWithoutOPperators(5, 7)
	fmt.Println(l)
}

func SumWithoutOPperators(a int,b int)int{
	m := make([]int, a, a)
	n := make([]int, b, b)
	m = append(m, n...)
	return len(m)
}