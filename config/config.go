package config

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	AppConfig()
}

func AppConfig() {
	viper.SetConfigName("restapi")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalln("🌐 Config has already been mastered globally!")
		} else {
			log.Fatalln("Error is : ", err)
		}
	}
}
