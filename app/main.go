package main

import (
	"app/app/wbhandler"
	"fmt"
	"log"
	"net/http"
	"os"
	"gopkg.in/yaml.v2"
	"strconv"
)


type Config struct {
	Service struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		TLS  struct {
			CertFile string `yaml:"cert_file"`
			KeyFile  string `yaml:"key_file"`
		} `yaml:"tls"`
	} `yaml:"service"`
    
}

func LoadConfig(file string) (*Config, error) {
	// Read the YAML configuration file and get config
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func main() {
	configFile := "app/config.yaml"
	conf , err := LoadConfig(configFile)
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	port :=  strconv.Itoa(conf.Service.Port)
	
	if len(os.Args) > 1 && os.Args[1] != "" {
		port = os.Args[1]
	}
	fmt.Println(conf.Service.Host+":"+port)
	http.HandleFunc("/webhook", wbhandler.WebHookHandler)
	log.Println("Starting HTTPS server on port " + port)
	//ListenAndServe starts an HTTPS server with a given address and handler.
	err2 := http.ListenAndServeTLS(conf.Service.Host+":"+port,
	 conf.Service.TLS.CertFile,conf.Service.TLS.KeyFile, nil)
	if err2 != nil {
		log.Fatal("Failed to start server:", err2)
	}

}
