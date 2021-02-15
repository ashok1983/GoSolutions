package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello world")
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	return
}
