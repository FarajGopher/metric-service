package config

import (
	"log"
    "github.com/spf13/viper"
)

type Config struct {
    Server struct {
        Port string
    }
    Auth struct {
        Username string
        Password string
    }
}

var AppConfig Config

func LoadConfig() {
    viper.SetConfigName("config")      
    viper.SetConfigType("yaml")         
    viper.AddConfigPath("./config") 

    // Read config file
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file: %v", err)
    }

    // Unmarshal into AppConfig
    if err := viper.Unmarshal(&AppConfig); err != nil {
        log.Fatalf("Unable to decode config into struct: %v", err)
    }
}
