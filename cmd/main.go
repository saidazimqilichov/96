package main

import (
	"https/handlers"
	"https/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/items", handlers.GetItems).Methods("GET")
	r.HandleFunc("/items", handlers.CreateItem).Methods("POST")
	r.HandleFunc("/items/{id}", handlers.GetItem).Methods("GET")
	r.HandleFunc("/items/{id}", handlers.UpdateItem).Methods("PUT")
	r.HandleFunc("/items/{id}", handlers.DeleteItem).Methods("DELETE")

	cfg := config.Load()

	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	go func() {
		log.Printf("Server starting on https://localhost:%s", cfg.Port)
		if err := srv.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	config.GracefulShutdown(srv)
}
