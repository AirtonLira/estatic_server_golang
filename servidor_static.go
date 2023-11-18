package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./estatic_page"))
	http.Handle("/", fs)
	log.Println("Listening on: 8030", nil)
	err := http.ListenAndServe(":8030", nil)
	if err != nil {
		log.Fatal(err)
	}
}
