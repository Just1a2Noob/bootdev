package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	server := http.Server{}

	server.Addr = ":8080"
	server.Handler = mux

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Server cannot listen : %s", err)
	}
}
