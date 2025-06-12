package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/YOUR-USER-OR-ORG-NAME/YOUR-REPO-NAME/internal/pkg/db"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")
	db.ConnectDatabase()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("welcome to the server kharcha kaha"))

	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	log.Printf("Server is running on http://localhost:%s", port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("could not start start server : %v", err)
	}
}
