package main

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/bentrevor/calhoun/app"
)

func main() {
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/echo", EchoHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	fmt.Println("(fmt.Println) server started on 8080")
	log.Print("(log.Print) server started on 8080")
}
