package main

import (
	"log"

	"github.com/matsumo-and/dr-score-gen/internal/infrastructure/router"
)

func main() {
	r := router.NewRouter()

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
