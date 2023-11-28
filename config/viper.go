package config

import (
	"os"

	"github.com/spf13/viper"
)

type EnvVars struct {
	MONGODB_URI    string `mapstructure:"MONGODB_URI"`
	DATABASE       string `mapstructure:"DATABASE"`
	PORT           string `mapstructure:"PORT"`
	AUTH0_DOMAIN   string `mapstructure:"AUTH0_DOMAIN"`
	AUTH0_AUDIENCE string `mapstructure:"AUTH0_AUDIENCE"`
	DVLA_API       string `mapstructure:"DVLA_API"`
	DVSA_API       string `mapstructure:"DVSA_API"`
}

func LoadConfig() (config EnvVars, err error) {
	env := os.Getenv("GO_ENV")
	if env == "production" {
		return EnvVars{
			MONGODB_URI:    os.Getenv("MONGODB_URI"),
			DATABASE:       os.Getenv("DATABASE"),
			PORT:           os.Getenv("PORT"),
			AUTH0_DOMAIN:   os.Getenv("AUTH0_DOMAIN"),
			AUTH0_AUDIENCE: os.Getenv("AUTH0_AUDIENCE"),
			DVLA_API:       os.Getenv("DVLA_API"),
			DVSA_API:       os.Getenv("DVSA_API"),
		}, nil
	}

	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
