package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	host := "0.0.0.0"
	port := 3000
	url := fmt.Sprintf("%s:%d", host, port)
	log.Printf("Listening on %s", url)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(".")))
	log.Fatal(http.ListenAndServe(url, mux))
}
