package main

import (
	"errors"
	"fmt"
	"os"
)

type AppConfig struct {
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
	DB_DSN      string

	Env string
	DSN string
}

func GetConfigFromEnv() (*AppConfig, error) {
	config := AppConfig{
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_DSN:      os.Getenv("DB_DSN"),
		Env:         os.Getenv("ENV"),
	}

	if config.DSN != "" {
		return &config, nil
	}

	if config.DB_HOST == "" ||
		config.DB_USER == "" ||
		config.DB_PASSWORD == "" ||
		config.DB_NAME == "" ||
		config.DB_PORT == "" {
		return nil, errors.New("Missing db config")
	}

	config.DSN = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DB_HOST, config.DB_USER, config.DB_PASSWORD, config.DB_NAME, config.DB_PORT)

	return &config, nil
}
