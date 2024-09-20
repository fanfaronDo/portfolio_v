package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

var defaultConfigPath = "./config/prod.yaml"

type Config struct {
	HttpServer `yaml:"http_server"`
	Mysql      `yaml:"mysql"`
}

type HttpServer struct {
	Address      string        `yaml:"address" env-default:"127.0.0.1"`
	Port         string        `yaml:"port" env-default:"8080"`
	Timeout      time.Duration `yaml:"timeout" env-default:"3s"`
	IdleTimeoute time.Duration `yaml:"idle_timeout" env-default:"30s"`
}

type Mysql struct {
	Address  string `yaml:"address" env-default:"127.0.0.1"`
	Port     string `yaml:"port" env-default:"3306"`
	User     string `yaml:"user" env-default:"root"`
	Password string `yaml:"password" env-default:"root"`
	Database string `yaml:"database" env-default:"portfolio"`
}

func ConfigLoad() (Config, error) {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = defaultConfigPath
		// log.Fatal("CONFIG_PATH environment variable not set")
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("CONFIG_PATH does not exist:", configPath)
	}

	var config Config

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatal("Cannot read config:", err)
		return Config{}, err
	}

	return config, nil
}
