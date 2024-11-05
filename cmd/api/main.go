package main

import (
	"go-challenge/cmd/api/server"
	"log"
)

func main() {
	r := server.New()
	if err := r.Run(); err != nil {
		log.Fatalf("error trying to run app %s", err.Error())
	}
}
