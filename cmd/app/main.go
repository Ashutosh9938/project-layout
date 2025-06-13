package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/TEST/NEW/internal/pkg/db"
	"github.com/TEST/NEW/internal/pkg/redis"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	_ = godotenv.Load(".env")
	db.ConnectDatabase()

	if err := redis.InitRedis(); err != nil {
		log.Fatalf(" Redis initialization failed: %v", err)
	}

	err := redis.Rdb.Set(redis.Ctx, "startup_key", "Redis is connected", 0).Err()
	if err != nil {
		log.Fatalf(" Failed to write test Redis key: %v", err)
	}
	val, err := redis.Rdb.Get(redis.Ctx, "startup_key").Result()
	if err != nil {
		log.Fatalf(" Failed to read test Redis key: %v", err)
	}

	log.Println(" Redis test value:", val)
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
