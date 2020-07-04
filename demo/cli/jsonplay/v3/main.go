package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("test for http server")
	server := NewPlayerServer(NewInMemoryStore())

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen to port 5000 %v", err)
	}
}
