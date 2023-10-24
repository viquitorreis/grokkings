package main

import "fmt"

func main() {
	numbers := []int{9231,
		337,
		9716,
		9309,
		9845,
		1919,
		3230,
		1589,
		7085,
		4665,
		7436,
		7202,
		3917,
		8930,
		2790,
		6558,
		6329,
		9871,
		3642,
		3613,
	}

	sortedNums := quicksort(numbers)

	fmt.Println(sortedNums)

}

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else {
		pivot := arr[0]
		less := []int{}
		greater := []int{}

		for _, num := range arr[1:] {
			if num < pivot {
				less = append(less, num)
				fmt.Println("less =>", less)
			} else {
				greater = append(greater, num)
				fmt.Println("greater =>", greater)
			}
		}
		fmt.Println("final less =>", less)
		fmt.Println("final greater =>", greater)
		less = append(quicksort(less), pivot)
		greater = quicksort(greater)

		return append(less, greater...)
	}

}
