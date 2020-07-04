package main

import "fmt"

func main() {
	fmt.Println("Happy Father's Day")

	var op string
	fmt.Scanln(&op)

	switch op {
	case "add":
		fmt.Println(add(10000, 20000))
	case "sub":
		fmt.Println(subtract(10000, 20000))
	case "mul":
		fmt.Println(multiply(10000, 20000))
	case "div":
		fmt.Println(divide(10000, 20000))
	default:
		fmt.Println("abrakadebra....")
	}
}

func add(x int, y int) int {
	return x + y
}

func subtract(x int, y int) int {
	return x - y
}

func multiply(x int, y int) int {
	return x * y
}

func divide(x int, y int) int {
	return x / y
}
