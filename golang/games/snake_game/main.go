package main

import (
	"log"

	"github.com/sachinjangam/programs/golang/games/snake_game/internal/storage"
)

func main() {
	store, err := storage.NewJsonStorage("data/scores.json")
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}
}
