package config

type (
	Config struct {
		Env             string `mapstructure:"ENV"`
		SwaggerUsername string `mapstructure:"SWAGGER_USERNAME"`
		SwaggerPassword string `mapstructure:"SWAGGER_PASSWORD"`
		Host            Host   `mapstructure:",squash"`
	}

	Host struct {
		Address      string `mapstructure:"HOST_ADDRESS"`
		Port         string `mapstructure:"HOST_PORT"`
		WriteTimeout int    `mapstructure:"HOST_WRITE_TIMEOUT"`
		ReadTimeout  int    `mapstructure:"HOST_READ_TIMEOUT"`
		IdleTimeout  int    `mapstructure:"HOST_IDLE_TIMEOUT"`
	}
)
