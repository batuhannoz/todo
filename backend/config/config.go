package config

import "os"

type Config struct {
	*MySQL
	*Server
}

type MySQL struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type Server struct {
	Port string
}

func GetConfig() *Config {
	return &Config{
		&MySQL{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_NAME"),
		},
		&Server{
			Port: os.Getenv("API_PORT"),
		},
	}
}
