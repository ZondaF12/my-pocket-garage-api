package config

import "github.com/spf13/viper"

type EnvVars struct {
	MONGODB_URI  string `mapstructure:"MONGODB_URI"`
	DATABASE string `mapstructure:"DATABASE"`
	PORT         string `mapstructure:"PORT"`
	AUTH0_DOMAIN   string `mapstructure:"AUTH0_DOMAIN"`
	AUTH0_AUDIENCE string `mapstructure:"AUTH0_AUDIENCE"`
	DVLA_API string `mapstructure:"DVLA_API"`
	DVSA_API string `mapstructure:"DVSA_API"`
}

func LoadConfig() (config EnvVars, err error) {
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