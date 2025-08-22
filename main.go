package main

import (
	"log"
	"net/http"
)

func main() {
	server := http.Server {
		Addr : "127.0.0.1:8080",
	}	
	if err := server.ListenAndServe(); err != nil {
		log.Println("Error starting server:", err)
	}
}