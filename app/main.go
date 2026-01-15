package main

import (
	"app/app/wbhandler"
	"log"
	"net/http"
	"os"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func main() {

	certPath := os.Getenv("TLS_CERT_PATH")
	keyPath := os.Getenv("TLS_KEY_PATH")
	if certPath == "" {
		certPath = "./certs/tls.crt"
	}
	if keyPath == "" {
		keyPath = "./certs/tls.key"
	}


	if !fileExists(certPath) || !fileExists(keyPath) {
		log.Fatal("TLS certificate or key file does not exist.")
	}

	host := os.Getenv("SERVICE_HOST")
	if host == "" {
		host = "0.0.0.0"
	}
	port := "8443"
	if len(os.Args) > 1 && os.Args[1] != "" {
		port = os.Args[1]
	}
	
	http.HandleFunc("/webhook", wbhandler.WebHookHandler)
	log.Println("Starting HTTPS server on port " + port)
	//ListenAndServe starts an HTTPS server with a given address and handler.
	err2 := http.ListenAndServeTLS(host+":"+port, certPath, keyPath, nil)
	if err2 != nil {
		log.Fatal("Failed to start server:", err2)
	}

}
