package main

import "fmt"

func vardiacfunc(params ...interface{}) {
	for _, element := range params {
		fmt.Println(element)
	}
}

func main() {
	vardiacfunc("one", "two", "three")
}
