package main

import (
	"fmt"
)

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
Example:
Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

func main(){
	nums : [4]{2, 7, 11, 15}
	target := 9
	solution(nums, target)
}

func solution(nums []int, target int){
	for i, item := range nums{
		if item<target{
			temp := nums[i:i+1]
		}
	}
}

