package utils

import (
	"log"

	"github.com/spf13/viper"
)

func GetCustomConf(key, defVal string) string {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("config file not found")
		log.Println("Default config has been loaded")
		return defVal
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}
