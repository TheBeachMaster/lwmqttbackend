package main

import (
	"log"
	"net/http"

	"com.thebeachmaster/mqttbackend/server"
)

func main() {
	log.Println("Starting Server...")
	port := ":8044"

	muxer := http.NewServeMux()

	server := server.NewServer(port, muxer)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
