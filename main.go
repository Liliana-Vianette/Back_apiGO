package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	server := http.ListenAndServe(":9091", router)
	log.Fatal(server)
}
