package config

import (
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(env string) {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing config file")
	}
}

func GetConfig() *viper.Viper {
	return config
}
