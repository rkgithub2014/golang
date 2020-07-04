package main

import "fmt"

//WriteOnly writes to channel as input order
func WriteOnly(channel chan<- int, order int) {
	channel <- order
}

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
