package main

import (
	"fmt"
	"time"
)

func tickFunc() {
	ticker := time.NewTicker(5000 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(16000 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

func main() {

	tickFunc()
}
