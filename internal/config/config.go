package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

/*
	Thank you u/jxstack on Reddit for their advice on
	good environment variable practices.

	https://www.reddit.com/r/golang/comments/1dzxah6/comment/lcjfw2h
*/

type Config struct {
	Logs LogConfig
	DB   PostgresConfig
	Host string
	Port string
}

type LogConfig struct {
	Style string
	Level string
}

type PostgresConfig struct {
	Username string
	Password string
	URL      string
	Port     string
}

/*
	This method returns an error for future-proofing.
	In the event that some part of the startup fails e.g. missing DB
	credentials, or log specs, we can return an appropriate error.

	Furthermore, this method should be called by the App constructor.
*/

func LoadConfig() (*Config, error) {
	cfg := &Config{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
		Logs: LogConfig{
			Style: os.Getenv("LOG_STYLE"),
			Level: os.Getenv("LOG_LEVEL"),
		},
		DB: PostgresConfig{
			Username: os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PWD"),
			URL:      os.Getenv("POSTGRES_URL"),
			Port:     os.Getenv("POSTGRES_PORT"),
		},
	}

	return cfg, nil
}
