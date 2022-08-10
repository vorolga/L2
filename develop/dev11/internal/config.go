package internal

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Host string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	return &Config{
		Port: os.Getenv("PORT"),
		Host: os.Getenv("HOST"),
	}, nil
}
