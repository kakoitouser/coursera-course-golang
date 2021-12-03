package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "pong")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
