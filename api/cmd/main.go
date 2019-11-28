package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alallema/picture_dictionnary.git/api/router"
)

func main() {
	router := router.NewRouter()
	fmt.Printf("Api Launch\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}
