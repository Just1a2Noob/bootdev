package main

import (
	"fmt"
	"net/http"
)

func main() {
	const filepathRoot = "."
	const port = "8080"
	mux := http.NewServeMux()
	server := http.Server{}

	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
	mux.HandleFunc("/healthz", handlerReadiness)

	server.Addr = ":" + port
	server.Handler = mux

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Server cannot listen : %s", err)
	}
}

func handlerReadiness(w http.ResponseWriter, request *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
