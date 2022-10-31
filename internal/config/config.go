package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

// Listen ...
type Listen struct {
	BindIP string `env:"BIND_IP" env-default:"0.0.0.0"`
	Port   string `env:"PORT" env-default:":9092"`
}

// AppConfig ...
type AppConfig struct {
	LogLevel string `env:"LOG_LEVEL" env-default:"trace"`
}

// Config ...
type Config struct {
	Listen        Listen
	AppConfig     AppConfig
	PostgresURL   string `env:"POSTGRES_URL"`
}

var instance *Config
var once sync.Once

// GetConfig ...
func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}

		if err := cleanenv.ReadEnv(instance); err != nil {
			helpText := "system"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
	})

	return instance
}