package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DB
}

type DB struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func New() *Config {
	initConfig()

	return &Config{
		DB: DB{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetString("DB_PORT"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			Database: viper.GetString("DB_NAME"),
		},
	}
}

func initConfig() {
	dir, _ := os.Getwd()
	fmt.Printf("\ncurrent dir: %+v\n", dir)
	dirs, _ := os.ReadDir(".")
	fmt.Printf("\nhere: %+v\n", dirs)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}
