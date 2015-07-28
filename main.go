package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/api/nodes", Nodes)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
