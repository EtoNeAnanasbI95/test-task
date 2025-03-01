package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"time"
)

type Config struct {
	Env                string        `env:"ENV" yaml:"env" env-default:"local" json:"env"`
	UserDB             string        `env:"USER_DB" yaml:"user_db" json:"userDb" env-required:"true"`
	HostDB             string        `env:"HOST_DB" yaml:"host_db" json:"hostDb" env-required:"true"`
	PortDB             int           `env:"PORT_DB" yaml:"port_db" json:"portDb" env-required:"true"`
	PasswordDB         string        `env:"PASSWORD_DB" yaml:"password_db" json:"passwordDb" env-required:"true"`
	NameDB             string        `env:"NAME_DB" yaml:"name_db" json:"nameDb" env-required:"true"`
	DBMS               string        `env:"DBMS" yaml:"dbms" json:"dbms" env-required:"true"`
	SslMode            string        `env:"SSL_MODE" yaml:"ssl_mode" json:"sslMode" env-default:"disable"`
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
