package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	uhHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, UrbanHive!")
	}

	http.HandleFunc("/", uhHandler)
	addr := ":" + os.Getenv("PORT")

	// addr := ":8080" use this port to test locally on your pc if you don't have the PORT env variable set

	log.Fatal(http.ListenAndServe(addr, nil))
}