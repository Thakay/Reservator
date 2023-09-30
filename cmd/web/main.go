package main

import (
	"fmt"
	"github.com/thakay/Reservator/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("starting appliaction on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
