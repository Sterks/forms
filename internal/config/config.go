package config

import (
	"fmt"
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

func Initial(configsDir string) (*Config, error) {
	var cfg *Config
	// viper.SetConfigFile("../../configs/main.yml")
	viper.SetConfigFile("configs/main.yml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	if err2 := viper.Unmarshal(&cfg); err2 != nil {
		logger.Errorf("Не могу сдалать Unmarshal для конфигурационного файла - %v", err2)
		return nil, err2
	}
	return cfg, nil
}
