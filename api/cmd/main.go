package main

import (
	"fmt"
	"log"

	"net/http"

	"../router"
)

func main() {
	router := router.NewRouter()
	fmt.Printf("Api Launch\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}
