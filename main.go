package main

import (
	"fmt"
)

//	Given an array of integers nums and an integer target,
//	return indices of the two numbers such that they add up to target.
//  TwoSum function returns the indices of the two numbers from a massive that add up to the target
func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for index := i + 1; index < len(nums); index++ {
			if nums[index] == target-nums[i] {
				return []int{i, index}
			}
		}
	}
	return nil
}


func main() {
	fmt.Printf("Start of TwoSum function \n")
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(TwoSum(nums, target))

	nums = []int{3, 5, 7, 9}
	target = 15
	fmt.Println(TwoSum(nums, target))

	nums = []int{-1, -3, -5, -8}
	target = -8
	fmt.Println(TwoSum(nums, target))

	fmt.Printf("End of TwoSum function \n \n")

}
