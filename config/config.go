package config

import (
	"os"

	"github.com/spf13/viper"
)

type DB struct {
	URI      string `json:"uri"`
	Database string `json:"database"`
	Timeout  int    `json:"timeout"`
}

type AMQP struct {
	URI      string `json:"uri"`
	Exchange string `json:"exchange"`
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
	AMQP   `json:"amqp"`
}

func Register(configFile string, configType string, mode string) error {
	baseDirectory, err := os.Getwd()
	if err != nil {
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
		AMQP: AMQP{
			URI:      viper.GetString("AMQP_URI"),
			Exchange: viper.GetString("AMQP_EXCHANGE"),
			Timeout:  viper.GetInt("AMQP_TIMEOUT"),
		},
	}

	return config
}
