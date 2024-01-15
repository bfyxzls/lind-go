package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8012", http.FileServer(http.Dir("."))))
}
