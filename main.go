package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ok") // to write response body
}

func main() {
	mux := http.NewServeMux() //router
	//kaam: -> konsa url --> konsa handler

	mux.HandleFunc("/health", healthHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Server starting on : 8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
