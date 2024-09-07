package env

import (
	"github.com/caarlos0/env/v8"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	DBDriver string `env:"DB_DRIVER" envDefault:"postgres"`
	DBSource string `env:"DB_SOURCE" envDefault:"user=postgres password=postgres dbname=user_service sslmode=disable"`
	Port     string `env:"APP_PORT" envDefault:"8080"`
}

func LoadConfig() *Config {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
