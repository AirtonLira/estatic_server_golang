package main

import (
	"fmt"
	"net/http"
)

// Simple example write word "Ola" in page.
func ola(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Ola\n")
}

// read all information about Header request.
func header(w http.ResponseWriter, req *http.Request) {
	for name, header := range req.Header {
		for _, c := range header {
			fmt.Fprintf(w, "%v: %v \n", name, c)
		}
	}
}

func main() {
	http.HandleFunc("/ola", ola)
	http.HandleFunc("/header", header)

	http.ListenAndServe(":8090", nil)
}
