package config

import (
	"log/slog"
	"os"

	"github.com/spf13/viper"
)

type DB struct {
	URI      string `json:"uri"`
	Database string `json:"database"`
	Timeout  int    `json:"timeout"`
}

type Server struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Config struct {
	Server `json:"server"`
	DB     `json:"db"`
}

func Register(configFile string, configType string, mode string) error {
	baseDirectory, err := os.Getwd()
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	viper.AddConfigPath(baseDirectory)
	viper.SetConfigName(configFile)
	viper.SetConfigType(configType)

	if mode == "release" {
		viper.AutomaticEnv()
	} else {
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}

	return nil
}

func GetConfig() Config {
	config := Config{
		Server: Server{
			Name: viper.GetString("SERVER_NAME"),
			Host: viper.GetString("SERVER_HOST"),
			Port: viper.GetInt("SERVER_PORT"),
		},
		DB: DB{
			URI:      viper.GetString("DB_URI"),
			Database: viper.GetString("DB_DATABASE"),
			Timeout:  viper.GetInt("DB_TIMEOUT"),
		},
	}

	return config
}
