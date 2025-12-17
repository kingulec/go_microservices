package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := "8443"
	if len(os.Args) > 1 && os.Args[1] != "" {
		port = os.Args[1]
	}
	log.Println("Starting HTTPS server on port " + port)
	err := http.ListenAndServeTLS(":"+port, "certs/tls.crt", "certs/tls.key", nil)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}

}
