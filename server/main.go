package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello, World!")

	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT must be set")
	}

	rounter := chi.NewRouter()

	rounter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// router.

	server := &http.Server{
		Handler: rounter,
		Addr:    ":" + port,
	}

	log.Printf("Server is running on port %s", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port is", port)
}
