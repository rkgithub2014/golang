package main

import (
	"fmt"
)

func main() {
	fmt.Println("Test")
	result := SumAll([]int{1, 2, 3, 4, 5}, []int{10, 20})
	for _, resultSum := range result {
		fmt.Println("Result sum", resultSum)
	}
}

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllV1(numberToSum ...[]int) []int {
	lengthOfNumbers := len(numberToSum)
	sums := make([]int, lengthOfNumbers)
	for i, numbersIndex := range numberToSum {
		sums[i] = Sum(numbersIndex)
	}
	return sums
}

func SumAll(NumbersToSum ...[]int) []int {
	var sum []int
	for _, numbers := range NumbersToSum {
		sum = append(sum, Sum(numbers))
	}
	return sum
}
