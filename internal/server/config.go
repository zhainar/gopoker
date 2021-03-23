package server

import (
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Addr string
}

func NewConfig(configPath string) (*ServerConfig, error) {
	err := godotenv.Load(configPath)

	if err != nil {
		return nil, err
	}

	return &ServerConfig{
		Addr: os.Getenv("TCP_SERVER_ADDR"),
	}, nil
}
