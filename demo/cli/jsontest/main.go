package main

import (
	"encoding/json"
	"fmt"
)

type ssm struct {
	
}
type Address struct {
	Street  string
	City    string
	ZipCode string
}
type User struct {
	Name string
	Addr Address
}

func main() {
	rkHomeAddr := &Address{Street: "123, East street", City: "Dublin", ZipCode: "94568"}
	rkUser := &User{Name: "Frank", Addr: *rkHomeAddr}
	newAddr := &Address{Street: "5721, Old Westbury way", City: "Dublin", ZipCode: "94568"}
	// copy new address

	b, err := json.Marshal(rkUser)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Println(string(b))
}
