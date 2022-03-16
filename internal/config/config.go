package config

import (
	"forms/pkg/logger"
	"time"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Mongo MongoConfig `mapstructure:"mongo"`
		HTTP  HTTPConfig
	}

	MongoConfig struct {
		URL      string `mapstructure:"url"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"db_name"`
		Port     string `mapstructure:"port"`
	}

	HTTPConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	}
)

func Init(configDir string) (*Config, error) {
	// viper.SetConfigFile("../../configs/main.yml")
	viper.SetConfigFile("/configs/main.yml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var conf *Config

	if err := viper.Unmarshal(&conf); err != nil {
		logger.Error(err)
	}
	return conf, err
}
