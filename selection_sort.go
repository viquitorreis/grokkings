package main

import "fmt"

func main() {
	array := []int{43, 88, 57, 52, 85, 86, 30, 62, 97, 72, 88, 95, 50, 13, 78, 78, 53, 88, 19, 68, 50, 97, 76, 7, 44, 16, 36, 5, 34, 17, 38, 37, 86, 92, 8, 92, 63, 10, 87, 63, 72, 14, 37, 18, 77, 86, 40, 99, 94, 51, 83, 27, 68, 84, 33, 75, 42, 39, 9, 38, 71, 26, 3, 75, 22, 26, 74, 70, 17, 62, 71, 97, 89, 39, 72, 93, 46, 77, 25, 70, 45, 55, 40, 28, 93, 67, 33, 38, 87, 56, 24, 39, 80, 54, 8, 80, 68, 87, 83, 14}

	selectionSort := selectionSort(array)
	fmt.Println(selectionSort)
}

func selectionSort(arr []int) []int {
	newArr := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		smallest := findSmallest(arr)
		newArr[i] = arr[smallest]
		arr = append(arr[:smallest], arr[smallest+1:]...)
	}

	return newArr
}

func findSmallest(arr []int) int {
	smallest := arr[0]  // armazena o menor valor
	smallest_index := 0 // indice do menor valor

	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallest_index = i
		}
	}

	return smallest_index
}
