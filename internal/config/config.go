package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env      string `yaml:"env" env-required:"true"`
	Server   `yaml:"server"`
	Database `yaml:"database"`
}

type Server struct {
	Host string `yaml:"host" env-required:"true"`
	Port string `yaml:"port" env-required:"true"`
}

type Database struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     int32  `yaml:"port" env-required:"true"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
	DB       string `yaml:"db" env-required:"true"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Print("CONFIG_PATH env was not provided \n Searching in config/local.yaml")
		configPath = "config/local.yaml"
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file not exitst in provided path: %s", configPath)
	}

	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Error while loading env: %s", err)
	}

	return &cfg
}
