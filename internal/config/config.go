package config

import (
	"errors"
	"os"
	"sync"
)

// use "github.com/ilyakaznacheev/cleanenv" to read yaml

// config = &Config{}
//err := cleanenv.ReadConfig("config.yaml", config)

type Config struct {
	Db DbConfig `yaml:"db"`
}

type DbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

var config *Config
var once sync.Once

func NewConfig() (*Config, error) {
	//errorStr := ""
	//once.Do(func() {
	//	config = &Config{}
	//	err := godotenv.Load(".env")
	//	if err != nil {
	//		errorStr += fmt.Sprintf("Config error raised- %s;", err.Error())
	//		return
	//	}
	//})
	//
	//if errorStr != "" {
	//	return &Config{}, errors.New(errorStr)
	//}

	config = &Config{}

	if config.Db.Host = os.Getenv("DB_HOST"); config.Db.Host == "" {
		return &Config{}, errors.New("DB_HOST is empty")
	}

	if config.Db.Port = os.Getenv("DB_PORT"); config.Db.Port == "" {
		return &Config{}, errors.New("DB_PORT is empty")
	}

	if config.Db.Username = os.Getenv("POSTGRES_USER"); config.Db.Username == "" {
		return &Config{}, errors.New("POSTGRES_USER is empty")
	}

	if config.Db.Password = os.Getenv("POSTGRES_PASSWORD"); config.Db.Password == "" {
		return &Config{}, errors.New("POSTGRES_PASSWORD is empty")
	}

	if config.Db.Name = os.Getenv("POSTGRES_DB"); config.Db.Name == "" {
		return &Config{}, errors.New("POSTGRES_DB is empty")
	}

	return config, nil
}
