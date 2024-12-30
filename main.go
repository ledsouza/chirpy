package main

import (
	"log"
	"net/http"
)

func main() {
	const port = ":8080"
	mux := http.NewServeMux()

	var srv = &http.Server{
		Addr:    port,
		Handler: mux,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("error in the server: %s", err)
	}
}
