package server

import (
	"net/http"

	"github.com/kiing-dom/dngi-airdrop/internal/handlers"
)

func New(addr string) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/health", handlers.Health)
	mux.HandleFunc("/upload", handlers.Upload)

	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}
