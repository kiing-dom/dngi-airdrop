package main

import (
	"log"

	"github.com/kiing-dom/dngi-airdrop/internal/server"
)

func main() {
	srv := server.New(":8000")

	log.Println("starting server on port :8000")

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
