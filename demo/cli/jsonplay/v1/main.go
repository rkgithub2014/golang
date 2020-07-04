package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("test")
	storeRef := NewPlayStore()

	fmt.Println("Store Reference", &storeRef, storeRef)
	st1 := NewStoreInstance()
	st1.RecordWin("kholi")
	st1.RecordWin("dhoni")
	st1.RecordWin("kholi")
	fmt.Println("Store Instance1", &st1, st1)

	st2 := NewStoreInstance()
	st2.RecordWin("sachin")
	st2.RecordWin("viru")
	st2.RecordWin("sachin")
	fmt.Println("Store Instance2", &st2, st2)

	server := &PlayerServer{NewPlayStore()}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}

}
