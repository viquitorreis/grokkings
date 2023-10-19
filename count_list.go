package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5}
	countItems := countList(list, 0)

	fmt.Println("A conta de todos os elementos da lista foi: ", countItems)

}

func countList(list []int, count int) int {

	if len(list) == 0 {
		return count
	}

	return countList(list[1:], count+1)
}
