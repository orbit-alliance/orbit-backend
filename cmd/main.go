package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar .env:", err)
	}

	cfg := load()

	app, cleanup := NewApp(cfg)
	defer cleanup()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := app.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
