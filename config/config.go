package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB          DB
	Server      Server
	ExternalAPI ExternalAPI
}

type Server struct {
	APIPathPrefix string
	Host          string
	Port          int
}

type ExternalAPI struct {
	EnrichUserDataEndpoint string
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
		Server: Server{
			APIPathPrefix: viper.GetString("server.apiPathPrefix"),
			Host:          viper.GetString("server.host"),
			Port:          viper.GetInt("server.port"),
		},
		ExternalAPI: ExternalAPI{
			EnrichUserDataEndpoint: viper.GetString("externalApi.enrichUserData.endpoint"),
		},
	}
}

func initConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}
