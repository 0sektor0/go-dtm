package main

import (
	"log"
)

func main() {
	server, err := NewServer()
	if(err != nil) {
		log.Panicln(err)
	}

	server.Start()
}