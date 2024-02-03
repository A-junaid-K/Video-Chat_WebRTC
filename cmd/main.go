package main

import (
	"log"

	"github.com/A-junaid-K/VideoChat_WebRTC/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalln(err.Error())
	}
	
}
