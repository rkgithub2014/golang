package main

import "fmt"

func main() {
	var ints []int
	channel := make(chan int, 10)

	for index := 0; index < 10; index++ {
		go WriteOnly(channel, index)
	}

	for i := range channel {
		ints = append(ints, i)

		if len(ints) == 10 {
			break
		}
	}
	fmt.Printf("Ints %v", ints)
}

// WriteOnly serves the purpose of demonstrating
// a method that writes to a write-only channel.
func WriteOnly(channel chan<- int, order int) {
	channel <- order
}
