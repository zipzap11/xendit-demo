package config

import (
	"log"

	"github.com/spf13/viper"
)

type XenditConfig struct {
	WriteKey string `mapstructure:"WRITE_KEY_XENDIT"`
	ReadKey  string `mapstructure:"READ_KEY_XENDIT"`
}

func LoadEnv(path string) (XenditConfig, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
		// return XenditConfig{}, err
	}

	var config XenditConfig

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal(err)
	}

	return config, nil
}
