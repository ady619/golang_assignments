package main

import (
    "fmt"
	"strings"
)

// Given an integer n, return an array ans of length n + 1 such that 
// for each i (0 <= i <= n), ans[i] is the number of 1's in the binary 
// representation of i.  
// Input: n = 2
// Output: [0,1,1]

func main(){
	ans := solution(100)
	fmt.Println(ans)
}

func solution(n int)([]int){
	ans := make([]int, n+1)
	for i:=0;i<n+1;i++{
		bin := fmt.Sprintf("%b", i)
		count := strings.Count(bin, "1")
		ans[i] = count
	}
	return ans
}