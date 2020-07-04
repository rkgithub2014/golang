package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//wgWaitGroup()
	//useChannel()
	//sendReceive()
	useSelect()
}

func useSelect() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {

		for {
			c1 <- "Every 500 msec"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every 2 seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}

	}
}

func sendReceive() {
	c := make(chan string, 2)
	c <- "hello"
	c <- "world"

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)

}

func useChannel() {
	c := make(chan string)
	go computeWithChannel("async", c)

	for msg := range c {
		fmt.Println(msg)
	}
}
func computeWithChannel(value string, c chan string) {
	for index := 0; index <= 5; index++ {
		c <- value
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}

func wgWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		compute(5)
		wg.Done()
	}()

	wg.Wait()

}

func compute(value int) {
	for index := 0; index < value; index++ {
		fmt.Println(index)
		time.Sleep(time.Second)

	}
}
