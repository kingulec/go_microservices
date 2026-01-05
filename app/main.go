package main

import (
	"log"
	"net/http"
	"os"
	"app/app/wbhandler"
)

func main() {
	port := "8443"
	if len(os.Args) > 1 && os.Args[1] != "" {
		port = os.Args[1]
	}
	http.HandleFunc("/webhook", wbhandler.WebHookHandler)
	log.Println("Starting HTTPS server on port " + port)
	// ListenAndServe starts an HTTPS server with a given address and handler.
	err := http.ListenAndServeTLS(":"+port, "certs/tls.crt", "certs/tls.key", nil)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}

}
