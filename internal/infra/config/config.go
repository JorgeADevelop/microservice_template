package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	AppConfig   AppConfig
	DBConfig    DBConfig
	AuthConfig  AuthConfig
	KafkaConfig KafkaConfig
}

type AppConfig struct {
	Host string
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type AuthConfig struct {
	Secret string
}

type KafkaConfig struct {
	Brokers []string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := &Config{
		AppConfig: AppConfig{
			Host: os.Getenv("APP_HOST"),
			Port: os.Getenv("APP_PORT"),
		},
		DBConfig: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_DATABASE"),
		},
		AuthConfig: AuthConfig{
			Secret: os.Getenv("AUTH_SECRET"),
		},
		KafkaConfig: KafkaConfig{
			Brokers: strings.Split(os.Getenv("KAFKA_BROKERS"), ","),
		},
	}

	return cfg
}
