package config

import (
	"os"

	"github.com/spf13/viper"
)

type Mongo struct {
	URI      string `json:"uri"`
	Database string `json:"database"`
	Timeout  int    `json:"timeout"`
}

type Rabbit struct {
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
	Mongo  `json:"mongo"`
	Rabbit `json:"rabbit"`
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
		Mongo: Mongo{
			URI:      viper.GetString("MONGO_URI"),
			Database: viper.GetString("MONGO_DATABASE"),
			Timeout:  viper.GetInt("MONGO_TIMEOUT"),
		},
		Rabbit: Rabbit{
			URI:      viper.GetString("RABBIT_URI"),
			Exchange: viper.GetString("RABBIT_EXCHANGE"),
			Timeout:  viper.GetInt("RABBIT_TIMEOUT"),
		},
	}

	return config
}
