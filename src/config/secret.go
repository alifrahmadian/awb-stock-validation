package config

import "github.com/spf13/viper"

func ViperBind() {
	viper.BindEnv("ENV")

	// Binding Swagger Auth
	viper.BindEnv("SWAGGER_USERNAME")
	viper.BindEnv("SWAGGER_PASSWORD")

	viper.BindEnv("HOST_ADDRESS")
	viper.BindEnv("HOST_PORT")
	viper.BindEnv("HOST_WRITE_TIMEOUT")
	viper.BindEnv("HOST_READ_TIMEOUT")
	viper.BindEnv("HOST_IDLE_TIMEOUT")
}
