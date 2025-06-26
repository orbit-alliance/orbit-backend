package main

import (
	"log"

	env "github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	MONGO_URL          string `env:"MONGO_URL"`
	APP_PORT           string `env:"APP_PORT"   envDefault:"8080"`
	PROJECT_START_DATE string `env:"PROJECT_START_DATE" envDefault:"2025-02-03T00:00:00Z"`
	CLIENT_ID_42       string `env:"CLIENT_ID_42"`
	CLIENT_SECRET_42   string `env:"CLIENT_SECRET_42"`
	REDIRECT_URI_42    string `env:"REDIRECT_URI_42" envDefault:"http://localhost:3000/app/wallet"`
	ETH_RPC_URL        string `env:"ETH_RPC"`
	ETH_PRIVATE_KEY    string `env:"ETH_PRIVATE_KEY"`
	CONTRACT_ADDRESS   string `env:"CONTRACT_ADDRESS"`
	ALLOWED_ORIGINS    string `env:"CORS_ORIGINS" envDefault:"*"`
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Print(".env not found")
	}
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	return cfg
}
