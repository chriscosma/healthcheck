package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("ok"))
		if err != nil {
			log.Printf("error while writing ok response: %s", err)
		}
	})

	port := os.Getenv("HEALTHCHECK_PORT")
	if port == "" {
		port = "9999"
	}

	host := fmt.Sprintf(":%s", port)

	log.Printf("listening on %s", host)
	err := http.ListenAndServe(host, nil)
	if err != nil {
		log.Fatalf("error while serving: %s", err)
	}
}
