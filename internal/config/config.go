package config

import "github.com/spf13/viper"

type (
	Config struct {
		Mongo MongoConfig `mapstructure:"mongo"`
	}

	MongoConfig struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		DBName   string `mapstructure:"db_name"`
		Port     string `mapstructure:"port"`
	}
)

func Init(configDir string) (config *Config, err error) {
	viper.AddConfigPath(configDir)
	viper.SetConfigFile("main.yml")
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
}
