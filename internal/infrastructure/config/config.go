package config

import (
	"os"
	"sync"
)

type Config struct {
	App   Application
	Mysql Database
}

type Application struct {
	Name        string
	Environment string
	Port        string
}

type Database struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

type BasicService struct {
	Host string
	Port string
}

var (
	instance *Config
	once     sync.Once
)

func Load() *Config {
	once.Do(func() {
		instance = &Config{
			App: Application{
				Name:        os.Getenv("APP_NAME"),
				Environment: os.Getenv("APP_ENV"),
				Port:        os.Getenv("APP_PORT"),
			},
			Mysql: Database{
				Host:     os.Getenv("DB_HOST"),
				Port:     os.Getenv("DB_PORT"),
				Name:     os.Getenv("DB_NAME"),
				User:     os.Getenv("DB_USER"),
				Password: os.Getenv("DB_PASSWORD"),
			},
		}
	})

	return instance
}
