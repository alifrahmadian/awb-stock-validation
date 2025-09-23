package config

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func LoadConfig() (config Config, err error) {
	viper.AutomaticEnv()

	// Check if the .env file exists
	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		viper.SetConfigType("env")

		if err := viper.ReadInConfig(); err != nil {
			logrus.Error("error load .env file")
			return config, err
		}
	}

	// do viper bind
	ViperBind()

	err = viper.Unmarshal(&config)

	return config, err
}
