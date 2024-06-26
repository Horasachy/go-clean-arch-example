package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	MongoUrl string `mapstructure:"MONGO_URL"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
