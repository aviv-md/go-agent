package config

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

type Environment struct {
	AI
}

type AI struct {
	BaseURL string
	APIKey  string
}

var Env *Environment = nil

func Load() *Environment {
	if Env != nil {
		log.Debug("Environment already loaded")
		return Env
	}

	err := godotenv.Load()
	if err != nil {
		log.Warn("Error loading .env file")
	}

	Env = &Environment{
		AI: AI{
			BaseURL: os.Getenv("AI_BASE_URL"),
			APIKey:  os.Getenv("AI_API_KEY"),
		},
	}

	if Env.BaseURL == "" {
		log.Fatal("AI_BASE_URL is required")
	}

	if Env.APIKey == "" {
		log.Fatal("AI_API_KEY is required")
	}

	return Env
}
