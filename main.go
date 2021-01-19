package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
	  log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	

	uhHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, UrbanHive!")
	}

	http.HandleFunc("/", uhHandler)
	addr := ":" + os.Getenv("PORT")

	// addr := ":8080" use this port to test locally on your pc if you don't have the PORT env variable set

	log.Fatal(http.ListenAndServe(addr, nil))
}