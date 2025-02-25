package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"time"
)

type Config struct {
	Env                string        `env:"ENV" yaml:"env" env-default:"local" json:"env"`
	ConnectionString   string        `env:"CONNECTION_STRING" yaml:"connection_string" json:"connection-string" env-required:"true"`
	ApiPort            int           `env:"API_PORT" yaml:"api_port" env-default:"8080" json:"port"`
	ApiTimeout         time.Duration `env:"API_TIMEOUT" yaml:"api_timeout" env-default:"30min" json:"timeout"`
	ExternalApiUrlBase string        `env:"EXTERNAL_API_URL_BASE" yaml:"external_api_url_base" env-required:"true" json:"external_api_url_base"`
}

func MustLoadConfig(envFilePath string) *Config {
	if err := godotenv.Load(envFilePath); err != nil {
		log.Fatalf("Error load .env file, err: %v", err)
	}
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatal("Error load env variables")
	}

	return &cfg
}
