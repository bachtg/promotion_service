package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		HTTPPort string `yaml:"httpport"`
	} `yaml:"server"`

	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

func LoadConfig() *Config {
	var cfg Config

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return &cfg
}
