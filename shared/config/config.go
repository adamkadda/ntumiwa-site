package config

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

/*
	Thank you u/jxstack on Reddit for their advice on
	good environment variable practices.

	https://www.reddit.com/r/golang/comments/1dzxah6/comment/lcjfw2h
*/

type Config struct {
	Logs       LogConfig
	DB         PostgresConfig
	HTTPClient HTTPClientConfig
	Host       string
	Port       string
	ServerType string
}

type LogConfig struct {
	Style string
	Level string
}

type PostgresConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

type HTTPClientConfig struct {
	MaxIdleConns        int
	MaxIdleConnsPerHost int
	IdleConnTimeoutSec  int
	RequestTimeoutSec   int
}

func (c *HTTPClientConfig) NewHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        c.MaxIdleConns,
			MaxIdleConnsPerHost: c.MaxIdleConnsPerHost,
			IdleConnTimeout:     time.Duration(c.IdleConnTimeoutSec) * time.Second,
		},
		Timeout: time.Duration(c.RequestTimeoutSec) * time.Second,
	}
}

func (pg *PostgresConfig) Connect() (*sql.DB, error) {
	if pg.DBName == "" {
		return nil, nil
	}

	connStr := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		pg.Username,
		pg.Password,
		pg.Host,
		pg.Port,
		pg.DBName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return db, nil
}

/*
	This method returns an error for future-proofing.
	In the event that some part of the startup fails e.g. missing DB
	credentials, or log specs, we can return an appropriate error.

	Furthermore, this method should be called by the App constructor.
*/

func LoadConfig() (*Config, error) {
	cfg := &Config{
		Host:       os.Getenv("HOST"),
		Port:       os.Getenv("PORT"),
		ServerType: os.Getenv("TYPE"),
		Logs: LogConfig{
			Style: os.Getenv("LOG_STYLE"),
			Level: os.Getenv("LOG_LEVEL"),
		},
		DB: PostgresConfig{
			DBName:   os.Getenv("POSTGRES_DB"),
			Username: os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PWD"),
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
		},
	}

	return cfg, nil
}
