package main

import "fmt"

func main() {
	nums := []int{5, 4, 3, 1, 2}

	sumNums := sum(nums)
	fmt.Println(sumNums)
}

func sum(nums []int) int {
	if len(nums) == 0 { // base case
		return 0
	}
	return nums[0] + sum(nums[1:])
}
