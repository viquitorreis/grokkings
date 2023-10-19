package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5, 10, 6}
	maxNum := maxNum(list, 0)

	fmt.Println("O maior número da lista é: ", maxNum)

}

func maxNum(list []int, max int) int {

	// base case
	if len(list) == 0 {
		return max
	}

	if list[0] > max {
		max = list[0]
	}

	return maxNum(list[1:], max)
}
