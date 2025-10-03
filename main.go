package main

import (
	"log"
	"net/http"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	mux := http.NewServeMux() // Mux server

	fs := http.StripPrefix("/app", http.FileServer(http.Dir("."))) // FileServer, stripped prefix
	mux.Handle("/app/", fs)                                        // Filserver handler
	mux.HandleFunc("/healthz", healthz)

	server := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}
	log.Println("Server starting on :8080") // log error
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
