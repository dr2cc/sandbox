package main

import "fmt"

func main() {
	fmt.Println(maxnumber([]int{2, 1, 6, 3, 8, 32, 16}))
	fmt.Println(maxindex([]int{2, 1, 6, 3, 8, 32, 16}))
}

func maxnumber(numbers []int) int {
	var max int

	for _, number := range numbers {
		if number > max {
			max = number
		}

	}
	return max
}

func maxindex(numbers []int) int {
	var max int
	var index int

	for i, number := range numbers {
		if number > max {
			max = number
			index = i
		}

	}
	return index
}
